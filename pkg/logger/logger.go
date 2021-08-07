package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func init() {
	cfg := zap.NewDevelopmentConfig()
	// cfg := zap.NewProductionConfig()
	// cfg.OutputPaths = []string{"stdout", "./out/app.log"}
	cfg.OutputPaths = []string{"stderr", "./out/app.log"}
	Logger, _ = cfg.Build()

	zap.ReplaceGlobals(Logger)
	Sugar = Logger.Sugar()
}
