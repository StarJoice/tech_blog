package test

import (
	"github.com/StarJoice/tools/ginx/gctx"
	"github.com/StarJoice/tools/ginx/session"
)

func Init() {
	// 本地session，测试使用
	session.SetDefaultProvider(&SessionProvider{})
}

type SessionProvider struct {
}

func (s SessionProvider) NewSession(ctx *gctx.Context, uid int64, jwtData map[string]string, sessData map[string]any) (session.Session, error) {
	return nil, nil
}

func (s SessionProvider) Get(ctx *gctx.Context) (session.Session, error) {
	val, _ := ctx.Get("_session")
	return val.(session.Session), nil
}

func (s SessionProvider) UpdateClaims(ctx *gctx.Context, claims session.Claims) error {
	//TODO implement me
	panic("implement me")
}

func (s SessionProvider) RenewAccessToken(ctx *gctx.Context) error {
	//TODO implement me
	panic("implement me")
}
