package gorms

import (
	"context"
	"errors"
	"time"

	"github.com/xuender/oils/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type gormLogger struct {
	level                     logger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func NewGormLogger(slows ...time.Duration) logger.Interface {
	var slow time.Duration
	if len(slows) > 0 {
		slow = slows[0]
	} else {
		slow = 500 * time.Millisecond // nolint
	}

	return &gormLogger{
		level:         logger.Info,
		SlowThreshold: slow,
	}
}

func (p *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	p.level = level

	return p
}

func (p *gormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if p.level < logger.Info {
		return
	}

	logs.Debugf(str, args...)
}

func (p *gormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if p.level < logger.Warn {
		return
	}

	logs.Warnf(str, args...)
}

func (p *gormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if p.level < logger.Error {
		return
	}

	logs.Errorf(str, args...)
}

func (p *gormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	yield func() (sql string, rowsAffected int64),
	err error,
) {
	if p.level <= 0 {
		return
	}

	elapsed := time.Since(begin)

	switch {
	case err != nil && p.level >= logger.Error &&
		(!p.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := yield()
		logs.Errorw(
			"trace",
			"err", err,
			"elapsed", elapsed,
			"rows", rows,
			"sql", sql,
		)
	case p.SlowThreshold != 0 && elapsed > p.SlowThreshold && p.level >= logger.Warn:
		sql, rows := yield()
		logs.Warnw("trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	case p.level >= logger.Info:
		sql, rows := yield()
		logs.Debugw("trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	}
}
