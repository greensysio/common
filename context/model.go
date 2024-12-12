package context

import (
	modelSE "bitbucket.org/greensys-tech/security/model"
	"context"
	"github.com/labstack/echo/v4"
	"time"
)

type (
	CustomContext struct {
		EchoContext echo.Context
		Timeout     time.Duration
		Locale      string
		TokenInfo   modelSE.TokenInfo
		TokenStr    string
	}
)

// ------- Get funcs -------
func (ctx *CustomContext) GetEchoContext() echo.Context {
	return ctx.EchoContext
}
func (ctx *CustomContext) GetTimeout() time.Duration {
	return ctx.Timeout
}
func (ctx *CustomContext) GetLocale() string {
	return ctx.Locale
}
func (ctx *CustomContext) GetTokenInfo() modelSE.TokenInfo {
	return ctx.TokenInfo
}
func (ctx *CustomContext) GetTokenStr() string {
	return ctx.TokenStr
}
func (ctx *CustomContext) GetContext() context.Context {
	if ctx.EchoContext.Request().Context().Err() != nil {
		tempCtx := context.WithValue(context.Background(), echo.HeaderXRequestID, ctx.EchoContext.Request().Header.Get(echo.HeaderXRequestID))
		tempCtx, _ = context.WithTimeout(tempCtx, ctx.GetTimeout())
		if len(ctx.EchoContext.Request().Header.Get("Authorization")) > 0 {
			tempCtx = context.WithValue(tempCtx, "Authorization", "Bearer "+ctx.EchoContext.Request().Header.Get("Authorization"))
		}
		if len(ctx.EchoContext.Request().Header.Get("Token-Internal")) > 0 {
			tempCtx = context.WithValue(tempCtx, "Token-Internal", ctx.EchoContext.Request().Header.Get("Token-Internal"))
		}
		return tempCtx
	}
	return ctx.EchoContext.Request().Context()
}

// ------- Set funcs -------
func (ctx *CustomContext) SetTimeout(timeout time.Duration) {
	ctx.Timeout = timeout
}
func (ctx *CustomContext) SetLocale(locale string) {
	ctx.Locale = locale
}
func (ctx *CustomContext) SetTokenInfo(tkInfo modelSE.TokenInfo) {
	ctx.TokenInfo = tkInfo
}
func (ctx *CustomContext) SetTokenStr(tkStr string) {
	ctx.TokenStr = tkStr
}
