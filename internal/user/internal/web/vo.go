package web

type signUpReq struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

type Profile struct {
	Id       int64  `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	AboutMe  string `json:"aboutme,omitempty"`
	Ctime    string `json:"ctime,omitempty"`
}

type editPasswordReq struct {
	//Email string `json:"email"`
	// 暂时还未引入验证码机制，所以暂时先定义出来，不使用
	//Code        int64  `json:"code"`
	OidPassword string `json:"oidpassword"`
	NewPassword string `json:"newpassword"`
}

// editReq 仅能更新用户信息下的非敏感字段（昵称、头像等等...），后续扩展再加入字段
type editReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	AboutMe  string `json:"aboutMe" binding:"required"`
}

type loginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
