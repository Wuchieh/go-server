package logger

import (
	"os"
	"time"

	"github.com/Wuchieh/go-server/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *Log

type Log struct {
	z *zap.Logger
}

func (l *Log) Sync() error {
	return l.z.Sync()
}

func (l *Log) Debug(args ...any) {
	l.z.Sugar().Debug(args...)
}

func (l *Log) Debugf(format string, args ...any) {
	l.z.Sugar().Debugf(format, args...)
}

func (l *Log) Info(args ...any) {
	l.z.Sugar().Info(args...)
}

func (l *Log) Infof(format string, args ...any) {
	l.z.Sugar().Infof(format, args...)
}

func (l *Log) Warn(args ...any) {
	l.z.Sugar().Warn(args...)
}

func (l *Log) Warnf(format string, args ...any) {
	l.z.Sugar().Warnf(format, args...)
}

func (l *Log) Error(args ...any) {
	l.z.Sugar().Error(args...)
}

func (l *Log) Errorf(format string, args ...any) {
	l.z.Sugar().Errorf(format, args...)
}

func (l *Log) Fatal(args ...any) {
	l.z.Sugar().Fatal(args...)
}

func (l *Log) Fatalln(args ...any) {
	l.z.Sugar().Fatalln(args...)
}

func (l *Log) Fatalf(format string, args ...any) {
	l.z.Sugar().Fatalf(format, args...)
}

func New(cfg config.Log) *Log {
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
	case config.LogFormatJSON:
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

	return &Log{z: zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))}
}

func Setup(cfg config.Log) {
	z := New(cfg)
	log = z
	zap.ReplaceGlobals(z.z)
}
