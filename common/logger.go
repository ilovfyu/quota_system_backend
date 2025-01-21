package common

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	g "quota_system/global"
	"time"
)

func getLogWriter() zapcore.WriteSyncer {
	format_date := time.Now().Format("2006-01-02")
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", g.ServiceConfig.LoggerConfig.SavePath, g.ServiceConfig.LoggerConfig.Prefix+format_date+".log"),
		MaxSize:    g.ServiceConfig.LoggerConfig.MaxSize,
		MaxBackups: g.ServiceConfig.LoggerConfig.MaxBackups,
		MaxAge:     g.ServiceConfig.LoggerConfig.MaxAge,
		Compress:   g.ServiceConfig.LoggerConfig.Compress,
	}
	return zapcore.AddSync(lumberjackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.InfoLevel)
	zapLogger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(zapLogger)
}
