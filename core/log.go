package core

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *zap.SugaredLogger

func initlog(logPath string, isColor bool, disableStdOut bool) *zap.SugaredLogger {
	var core zapcore.Core
	hook := lumberjack.Logger{Filename: logPath, MaxAge: 90}

	encConfigFile := zap.NewProductionEncoderConfig()
	encConfigOs := zap.NewProductionEncoderConfig()

	encConfigFile.EncodeTime = zapcore.ISO8601TimeEncoder
	encConfigOs.EncodeTime = zapcore.ISO8601TimeEncoder
	if !disableStdOut {
		if isColor {
			encConfigOs.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(encConfigFile), zapcore.AddSync(&hook), zap.DebugLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(encConfigOs), zapcore.AddSync(os.Stdout), zap.DebugLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(encConfigFile), zapcore.AddSync(&hook), zap.DebugLevel),
		)
	}

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func init() {
	SCST_GATEWAY_LOG := os.Getenv("SCST_GATEWAY_LOG")
	if _, err := os.Stat(SCST_GATEWAY_LOG); err != nil {
		os.MkdirAll(SCST_GATEWAY_LOG, 0755)
	}
	SCST_GATEWAY_LOG = "/var/log/scstgateway.log"
	l = initlog(SCST_GATEWAY_LOG, false, false)
	l.Infof("$SCST_GATEWAY_LOG:%s", SCST_GATEWAY_LOG)
}
