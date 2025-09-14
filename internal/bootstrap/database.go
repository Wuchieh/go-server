package bootstrap

import (
	"context"
	"time"

	orm "github.com/Wuchieh/go-server-orm"
	"github.com/Wuchieh/go-server/internal/config"
	"github.com/Wuchieh/go-server/internal/utils/logger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type log struct {
	gormLogger.Config
	level gormLogger.LogLevel
}

func (l log) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	l.level = level
	return l
}

func (l log) Info(_ context.Context, s string, i ...interface{}) {
	logger.GetLogger().Infof("[GORM] "+s, i...)
}

func (l log) Warn(_ context.Context, s string, i ...interface{}) {
	logger.GetLogger().Warnf("[GORM] "+s, i...)
}

func (l log) Error(_ context.Context, s string, i ...interface{}) {
	logger.GetLogger().Errorf("[GORM] "+s, i...)
}

func (l log) Trace(_ context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.level >= gormLogger.Error:
		logger.GetLogger().Errorf("[GORM] trace error: %v, elapsed: %v, rows: %d, sql: %s", err, elapsed, rows, sql)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.level >= gormLogger.Warn:
		logger.GetLogger().Warnf("[GORM] trace elapsed: %v, rows: %d, sql: %s", elapsed, rows, sql)
	case l.level == gormLogger.Info:
		logger.GetLogger().Infof("[GORM] trace elapsed: %v, rows: %d, sql: %s", elapsed, rows, sql)

	}
}

func databaseSetup() error {
	l := &log{
		Config: gormLogger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢查詢閾值
			IgnoreRecordNotFoundError: true,                   // 預設忽略 gorm.ErrRecordNotFound
			ParameterizedQueries:      false,                  // 可以改 true 來隱藏參數
		},
	}

	switch config.GetConfig().Log.Level {
	case config.LogLevelDebug:
		l.level = gormLogger.Info
	case config.LogLevelInfo:
		l.level = gormLogger.Info
	case config.LogLevelWarn:
		l.level = gormLogger.Warn
	case config.LogLevelError:
		l.level = gormLogger.Error
	default:
		l.level = gormLogger.Error // 預設至少輸出 error
	}

	cfg := gorm.Config{
		Logger:      l,
		PrepareStmt: true,
	}
	return orm.Setup(config.GetConfig().Database, &cfg)
}

func databaseClose() error {
	return orm.Close()
}
