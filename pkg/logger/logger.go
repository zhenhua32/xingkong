package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func init() {
	Logger, _ = zap.NewDevelopment()
	zap.ReplaceGlobals(Logger)
	Sugar = Logger.Sugar()
}
