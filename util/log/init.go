package log

import (
	"go-ent-demo/config"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.SugaredLogger

const (
	syncerFile    = "file"
	syncerConsole = "console"
)

func GetLogger() *zap.SugaredLogger {
	return logger
}

// InitLog init global logger
func InitLog() {
	cfg := config.GetConfig().Log

	cores := zapcore.NewTee(
		zapcore.NewCore(getEncoder(syncerConsole), getConsoleSyncer(), transportLevel(cfg.Level)),
		zapcore.NewCore(getEncoder(syncerFile), getFileSyncer(), transportLevel(cfg.Level)),
	)

	logger = zap.New(cores, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func getConsoleSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getFileSyncer() zapcore.WriteSyncer {
	cfg := config.GetConfig().Log

	fileWriter := &lumberjack.Logger{
		Filename:   cfg.File,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		LocalTime:  true,
		Compress:   cfg.Compress,
	}

	return zapcore.AddSync(fileWriter)
}

func getEncoder(syncer string) zapcore.Encoder {
	cfg := config.GetConfig().Log

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	switch syncer {
	case syncerConsole:
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case syncerFile:
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	default:
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	switch cfg.Format {
	case "json":
		return zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		return zapcore.NewConsoleEncoder(encoderConfig)
	default:
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
}

func transportLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
