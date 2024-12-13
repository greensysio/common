// Package log wrapped logrus functions
package log

import (
	"encoding/json"
	"runtime"

	"github.com/greensysio/common/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// -----------------------------------
// Logger uses for trace
// -----------------------------------

// ArgsCtx output message of print level
func ArgsCtx(ctx *context.CustomContext, mess string, args ...interface{}) {
	a, _ := json.Marshal(args)
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"args": string(a),
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Info(mess)
}

// PrintCtx output message of print level
func PrintCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Print(i...)
}

// PrintfCtx output format message of print level
func PrintfCtx(ctx *context.CustomContext, format string, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Printf(format, i...)
}

// DebugCtx output message of debug level
func DebugCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Debug(i...)
}

// Debugf output format message of debug level
func DebugfCtx(ctx *context.CustomContext, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Debugf(format, args...)
}

// InfoCtx output message of info level
func InfoCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Info(i...)
}

// InfofCtx output format message of info level
func InfofCtx(ctx *context.CustomContext, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Infof(format, args...)
}

// WarnCtx output message of warn level
func WarnCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Warn(i...)
}

// WarnfCtx output format message of warn level
func WarnfCtx(ctx *context.CustomContext, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Warnf(format, args...)
}

// ErrorCtx output message of error level
func ErrorCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Error(i...)
}

// ErrorfCtx output format message of error level
func ErrorfCtx(ctx *context.CustomContext, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Errorf(format, args...)
}

// FatalCtx output message of fatal level
func FatalCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Fatal(i...)
}

// FatalfCtx output format message of fatal level
func FatalfCtx(ctx *context.CustomContext, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Fatalf(format, args...)
}

// PanicCtx output message of panic level
func PanicCtx(ctx *context.CustomContext, i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Panic(i...)
}

// PanicfCtx output format message of panic level
func PanicfCtx(ctx *context.CustomContext, format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	singletonLogger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
		"id":   ctx.GetContext().Value(echo.HeaderXRequestID).(string),
	}).Panicf(format, args...)
}
