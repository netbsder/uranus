package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	config := GetAppConfig()
	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/%s.log", config.App.Name), // 日志文件路径
		MaxSize:    20,                                          // megabytes
		MaxBackups: 30,                                          // 最多保留30个备份
		MaxAge:     7,                                           // 最多保留7天
		Compress:   true,                                        // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&hook)

	// 设置日志级别 debug->info->warn->error
	var level zapcore.Level
	switch config.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	var core zapcore.Core
	if config.App.RunMode != "release" {
		core = zapcore.NewTee(
			// 打印在控制台
			zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.Lock(os.Stdout), level),
			// 打印在文件中
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level),
		)
	} else {
		// 打印在文件中
		core = zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), w, level)
	}

	logger = zap.New(core)
	logger.Debug("Logger init success!")
}

func GetLogger() *zap.Logger {
	return logger
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
