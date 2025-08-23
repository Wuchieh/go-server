package logger

import (
	"os"
	"time"

	"github.com/Wuchieh/go-server/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func Setup(cfg config.Log) {
	var zapConfig zap.Config

	switch cfg.Level {
	case config.LogLevelDebug:
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case config.LogLevelInfo:
		zapConfig = zap.NewProductionConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case config.LogLevelWarn:
		zapConfig = zap.NewProductionConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case config.LogLevelError:
		zapConfig = zap.NewProductionConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		zapConfig = zap.NewProductionConfig()
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// 時間格式
	zapConfig.EncoderConfig.TimeKey = "timestamp"
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// encoder
	var encoder zapcore.Encoder
	switch cfg.Format {
	case config.LogFormatJson:
		encoder = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	case config.LogFormatConsole:
		encoder = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
	}

	var cores []zapcore.Core

	// console 輸出
	if cfg.Console {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
			zapcore.AddSync(os.Stdout),
			zapConfig.Level,
		))
	}

	// info.log (只收 < ERROR)
	if cfg.OutputPath != "" {
		infoWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		})

		minLvl := zapConfig.Level.Level()
		cores = append(cores, zapcore.NewCore(
			encoder,
			infoWriter,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= minLvl && lvl < zapcore.ErrorLevel
			}),
		))
	}

	// error.log (只收 >= ERROR)
	if cfg.ErrorPath != "" {
		errorWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.ErrorPath, // 例如 logs/error.log
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		})

		cores = append(cores, zapcore.NewCore(
			encoder,
			errorWriter,
			zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.ErrorLevel
			}),
		))
	}

	// 合併 core
	core := zapcore.NewTee(cores...)
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	zap.ReplaceGlobals(log)
}

func GetLogger() *zap.Logger {
	return log
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Debugf(format string, a ...any) {
	log.Sugar().Debugf(format, a...)
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Infof(format string, a ...any) {
	log.Sugar().Infof(format, a...)
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Warnf(format string, a ...any) {
	log.Sugar().Warnf(format, a...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Errorf(format string, a ...any) {
	log.Sugar().Errorf(format, a...)
}

func Sync() {
	log.Sync()
}
