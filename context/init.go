package context

import (
	"context"
	"fmt"
	"time"

	authUtils "github.com/greensysio/security/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

// InitCustomCtx will return customed context and cancel func
func InitCustomCtx(c echo.Context, timeout time.Duration) (cusCtx *CustomContext, cancelFunc context.CancelFunc, errMsg string) {
	tempCtx := c.Request().Context()
	if tempCtx == nil {
		tempCtx = context.Background()
	}
	tempCtx = context.WithValue(tempCtx, echo.HeaderXRequestID, c.Request().Header.Get(echo.HeaderXRequestID))
	ctx, cancelFunc := context.WithTimeout(tempCtx, timeout)
	c.SetRequest(c.Request().WithContext(ctx))
	cusCtx = &CustomContext{
		EchoContext: c,
		Timeout:     timeout,
		Locale:      c.Request().Header.Get("Accept-Language"),
	}
	if len(cusCtx.Locale) == 0 {
		cusCtx.Locale = "VN"
	}
	if c.Get("user") != nil {
		cusCtx.TokenStr = c.Get("user").(*jwt.Token).Raw
		if cusCtx.TokenInfo, errMsg = authUtils.DecodeToken(c.Get("user")); len(errMsg) > 0 {
			errMsg = fmt.Sprintf("Error: %+v", errMsg)
		}
	}
	return
}

// InitCustomCtxFromOldOne will return new custom context
func InitNewCustomCtxFromOldOne(parent *CustomContext) (newCusCtx *CustomContext) {
	cusCtx := new(CustomContext)
	*cusCtx = *parent

	httpRq := parent.EchoContext.Request()
	ctx := context.WithValue(context.Background(), echo.HeaderXRequestID, parent.GetContext().Value(echo.HeaderXRequestID))
	if parent.GetContext().Value("Authorization") != nil {
		ctx = context.WithValue(ctx, "Authorization", "Bearer "+parent.GetContext().Value("Authorization").(string))
	}
	if parent.GetContext().Value("Token-Internal") != nil {
		ctx = context.WithValue(ctx, "Token-Internal", parent.GetContext().Value("Token-Internal").(string))
	}
	ctx, _ = context.WithTimeout(ctx, parent.GetTimeout())
	httpRq.WithContext(ctx)
	cusCtx.EchoContext.SetRequest(httpRq)
	return cusCtx
}

// InitNewCtxFromCustomCtx will return context and cancel func
func InitNewCtxFromCustomCtx(parent *CustomContext) (ctx context.Context, cancelFunc context.CancelFunc) {
	ctx = context.WithValue(context.Background(), echo.HeaderXRequestID, parent.GetContext().Value(echo.HeaderXRequestID))
	ctx, cancelFunc = context.WithTimeout(ctx, parent.GetTimeout())
	return
}

// RequestID middleware adds request ID
func RequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		randomString := random.String(32)
		if c.Request().Header.Get(echo.HeaderXRequestID) != "" {
			randomString = c.Request().Header.Get(echo.HeaderXRequestID) + "-" + randomString
		}
		c.Request().Header.Set(echo.HeaderXRequestID, randomString)
		return next(c)
	}
}
