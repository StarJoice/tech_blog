package web

import (
	"errors"
	"fmt"
	_ "github.com/StarJoice/tech_blog/docs"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tools/ginx/session"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
	"time"
)

var (
	emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	//emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	ErrDuplicateEmail    = service.ErrDuplicateEmail
	ErrInvalidPassword   = service.ErrInvalidPassword
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            service.UserService
	logger         *elog.Component
}

func NewUserHandle(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc:            svc,
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		logger:         elog.DefaultLogger,
	}
}

// PublicRoutes 公开路由
func (h *UserHandler) PublicRoutes(server *gin.Engine) {
	server.POST("/signup", ginx.WithRequest[signUpReq](h.SignUp))
	server.POST("/login", ginx.WithRequest[loginReq](h.Login))
	//server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
func (h *UserHandler) PrivateRoutes(server *gin.Engine) {
	user := server.Group("/user")
	user.GET("/profile", ginx.WithSession(h.Profile))
	//patch 在restful api 中的语义是更新部分资源，
	//对比而言，post的语义就更像是完整的资源更新，
	//所以此处使用patch来表示更新用户信息或者用户自己的密码，都是user表资源的一部分
	user.PATCH("/profile", ginx.WithSessionAndRequest[editReq](h.Edit))
	user.PATCH("/password", ginx.WithSessionAndRequest[editPasswordReq](h.UpdatePassword))
	// 还有更多路由例如：更换绑定邮箱等等的操作，暂时搁置，因为需要耦合email的服务，等实现后再来设计并实现
}

func (h *UserHandler) SignUp(ctx *ginx.Context, req signUpReq) (ginx.Result, error) {
	// 判断邮箱格式是否正确
	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		return systemErrorResult, fmt.Errorf("email regexp error: %w", err)
	}
	if !isEmail {
		return ginx.Result{Msg: "邮箱格式不正确"}, fmt.Errorf("邮箱:%s格式不正确", req.Email)
	}
	// 判断输入密码是否一致
	if req.Password != req.ConfirmPassword {
		return ginx.Result{Msg: "两次密码输入不一致"}, fmt.Errorf("用户输入密码不一致")
	}
	// 判断密码是否符合格式
	isPassword, err := h.passwordRexExp.MatchString(req.Password)
	if err != nil {
		return systemErrorResult, fmt.Errorf("password regexp error: %w", err)
	}
	if !isPassword {
		return ginx.Result{
			Msg: "密码必须包含至少一个字母、一个数字和一个特殊字符，并且长度至少为8个字符",
		}, fmt.Errorf("密码不符合格式要求")
	}
	// 调用下层service
	err = h.svc.Signup(ctx.Request.Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	switch {
	case err == nil:
		return ginx.Result{Msg: "注册成功"}, nil
	case errors.Is(err, ErrDuplicateEmail):
		return ginx.Result{Msg: "邮箱已注册，请返回登录页面登录"}, ErrDuplicateEmail
	default:
		return systemErrorResult, err
	}
}

func (h *UserHandler) Login(ctx *ginx.Context, req loginReq) (ginx.Result, error) {
	user, err := h.svc.Login(ctx.Request.Context(), req.Email, req.Password)
	switch {
	case err == nil:
		// 登录成功设置session
		_, err = session.NewSessionBuilder(ctx, user.Id).Build()
		if err != nil {
			return systemErrorResult, err
		}
		return ginx.Result{Msg: "登录成功"}, nil
	case errors.Is(err, ErrDuplicateEmail):
		return ginx.Result{Msg: "请检查邮箱或者密码"}, ErrDuplicateEmail
	default:
		return systemErrorResult, err
	}
}

func (h *UserHandler) Profile(ctx *ginx.Context, sess session.Session) (ginx.Result, error) {
	// 直接从session中拿到uid
	uid := sess.Claims().Uid
	u, err := h.svc.Profile(ctx.Request.Context(), uid)
	if err != nil {
		return ginx.Result{}, err
	}
	return ginx.Result{Data: Profile{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		AboutMe:  u.AboutMe,
		Ctime:    u.Ctime.Format(time.DateTime),
	}}, nil
}

func (h *UserHandler) Edit(ctx *ginx.Context, req editReq, sess session.Session) (ginx.Result, error) {
	uid := sess.Claims().Uid
	err := h.svc.UpdateNonSensitiveInfo(ctx.Request.Context(), domain.User{
		Id:       uid,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		AboutMe:  req.AboutMe,
	})
	if err != nil {
		return systemErrorResult, err
	}
	return ginx.Result{
		Msg: "用户信息更新完成",
	}, nil
}

func (h *UserHandler) UpdatePassword(ctx *ginx.Context, req editPasswordReq, sess session.Session) (ginx.Result, error) {
	// 检查新密码是否符合已有的密码规则
	isPassword, err := h.passwordRexExp.MatchString(req.NewPassword)
	if err != nil {
		return systemErrorResult, err
	}
	if !isPassword {
		return DataErrorResult, err
	}
	uid := sess.Claims().Uid
	err = h.svc.UpdatePassword(ctx.Request.Context(), uid, req.OidPassword, req.NewPassword)
	switch {
	case err == nil:
		return ginx.Result{Msg: "密码已更新"}, err
	case errors.Is(err, ErrInvalidPassword):
		return ginx.Result{Msg: "输入的原密码错误，请检查后重新输入"}, ErrInvalidPassword
	default:
		return systemErrorResult, err
	}
}
