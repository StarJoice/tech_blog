//@Date 2024/12/8 19:32
//@Desc

package web

type Profile struct {
	Id       int64  `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	AboutMe  string `json:"aboutme,omitempty"`
}
