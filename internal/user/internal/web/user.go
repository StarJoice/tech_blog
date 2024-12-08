//@Date 2024/12/5 00:59
//@Desc

package web

import (
	"errors"
	"fmt"
	"github.com/StarJoice/tech_blog/internal/user/internal/domain"
	"github.com/StarJoice/tech_blog/internal/user/internal/service"
	"github.com/StarJoice/tools/ginx/session"
	ginx "github.com/StarJoice/tools/ginx/wrapper"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/elog"
)

var (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	ErrDuplicateEmail    = service.ErrDuplicateEmail
)

type UserHandler struct {
	emailRexExp    *regexp.Regexp
	passwordRexExp *regexp.Regexp
	svc            service.UserService
	logger         *elog.Component
}

// PublicRoutes 公开路由
func (h *UserHandler) PublicRoutes(server *gin.Engine) {
	server.POST("/signup", ginx.WithRequest[signUpReq](h.SignUp))
	server.POST("/login", ginx.WithRequest[loginReq](h.Login))
}
func (h *UserHandler) PrivateRoutes(server *gin.Engine) {
	user := server.Group("/user")
	user.GET("/profile", ginx.WithSession(h.Profile))
	user.PUT("/profile", ginx.WithSessionAndRequest[editReq](h.Edit))
}

type signUpReq struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

func (h *UserHandler) SignUp(ctx *ginx.Context, req signUpReq) (ginx.Result, error) {
	//h.logger.Info("SignUp", elog.Any("email", req.Email))
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
	err = h.svc.Signup(ctx, domain.User{
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

type loginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(ctx *ginx.Context, req loginReq) (ginx.Result, error) {
	user, err := h.svc.Login(ctx, req.Email, req.Password)
	switch {
	// 登录成功设置session
	case err == nil:
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
	u, err := h.svc.Profile(ctx, uid)
	if err != nil {
		return ginx.Result{}, err
	}
	return ginx.Result{Data: Profile{
		Id:       u.Id,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		AboutMe:  u.AboutMe,
	}}, nil
}

// editReq 仅能更新用户信息下的非敏感字段（昵称、头像等等...），后续扩展再加入字段
type editReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	AboutMe  string `json:"aboutMe" binding:"required"`
}

func (h *UserHandler) Edit(ctx *ginx.Context, req editReq, sess session.Session) (ginx.Result, error) {
	uid := sess.Claims().Uid
	err := h.svc.UpdateNonSensitiveInfo(ctx, domain.User{
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

func NewUserHandle(svc service.UserService) *UserHandler {
	return &UserHandler{
		svc:            svc,
		emailRexExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRexExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		logger:         elog.DefaultLogger,
	}
}
