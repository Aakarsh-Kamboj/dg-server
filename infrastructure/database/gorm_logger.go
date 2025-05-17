package database

import (
	"context"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type ZapGormLogger struct {
	zap   *zap.Logger
	level gormlogger.LogLevel
}

func NewZapGormLogger(zapLogger *zap.Logger) gormlogger.Interface {
	return &ZapGormLogger{zap: zapLogger, level: gormlogger.Info}
}

func (l *ZapGormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	cloned := *l
	cloned.level = level
	return &cloned
}

func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Info {
		l.zap.Sugar().Infow(msg, data...)
	}
}

func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Warn {
		l.zap.Sugar().Warnw(msg, data...)
	}
}

func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Error {
		l.zap.Sugar().Errorw(msg, data...)
	}
}

func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level == gormlogger.Silent {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	switch {
	case err != nil && l.level >= gormlogger.Error:
		l.zap.Sugar().Errorw("SQL error", "error", err, "elapsed", elapsed, "rows", rows, "sql", sql)
	case elapsed > 200*time.Millisecond && l.level >= gormlogger.Warn:
		l.zap.Sugar().Warnw("Slow SQL", "elapsed", elapsed, "rows", rows, "sql", sql)
	case l.level >= gormlogger.Info:
		l.zap.Sugar().Infow("SQL", "elapsed", elapsed, "rows", rows, "sql", sql)
	}
}
