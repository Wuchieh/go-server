package bootstrap

import (
	"context"
	"time"

	orm "github.com/Wuchieh/go-server-orm"
	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/utils/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type ZapGormLogger struct {
	gormLogger.Config
	zapLogger *zap.Logger
}

func NewZapGormLogger(zapLogger *zap.Logger) *ZapGormLogger {
	return &ZapGormLogger{
		Config: gormLogger.Config{
			SlowThreshold: time.Second, // 超過多久算慢查詢
			LogLevel: func() gormLogger.LogLevel {
				switch config.GetConfig().Log.Level {
				case config.LogLevelInfo,
					config.LogLevelDebug:
					return gormLogger.Info
				case config.LogLevelWarn:
					return gormLogger.Warn
				case config.LogLevelError:
					return gormLogger.Error
				default:
					return gormLogger.Silent
				}
			}(), // 記錄等級
			IgnoreRecordNotFoundError: true, // 忽略 gorm.ErrRecordNotFound
			Colorful:                  false,
		},
		zapLogger: zapLogger,
	}
}

func (l *ZapGormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *ZapGormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.zapLogger.Sugar().Infof(msg, data...)
	}
}

func (l *ZapGormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.zapLogger.Sugar().Warnf(msg, data...)
	}
}

func (l *ZapGormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.zapLogger.Sugar().Errorf(msg, data...)
	}
}

func (l *ZapGormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.LogLevel >= gormLogger.Error:
		l.zapLogger.Error("sql error",
			zap.Error(err),
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
		)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLogger.Warn:
		l.zapLogger.Warn("slow sql",
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
		)
	case l.LogLevel == gormLogger.Info:
		l.zapLogger.Info("sql",
			zap.Duration("elapsed", elapsed),
			zap.String("sql", sql),
			zap.Int64("rows", rows),
		)
	}
}

func databaseSetup() error {
	cfg := gorm.Config{
		Logger: NewZapGormLogger(logger.GetLogger()),
	}
	return orm.Setup(config.GetConfig().Database, &cfg)
}

func databaseClose() error {
	return orm.Close()
}
