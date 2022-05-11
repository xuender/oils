package groms

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/xuender/oils/logs"
	"go.uber.org/zap"
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

	p.logger().Debugf(str, args...)
}

func (p *gormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if p.level < logger.Warn {
		return
	}

	p.logger().Warnf(str, args...)
}

func (p *gormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if p.level < logger.Error {
		return
	}

	p.logger().Errorf(str, args...)
}

func (p *gormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	call func() (sql string, rowsAffected int64),
	err error,
) {
	if p.level <= 0 {
		return
	}

	elapsed := time.Since(begin)

	switch {
	case err != nil && p.level >= logger.Error &&
		(!p.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := call()
		p.logger().Errorw(
			"trace",
			"err", err,
			"elapsed", elapsed,
			"rows", rows,
			"sql", sql,
		)
	case p.SlowThreshold != 0 && elapsed > p.SlowThreshold && p.level >= logger.Warn:
		sql, rows := call()
		p.logger().Warnw("trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	case p.level >= logger.Info:
		sql, rows := call()
		p.logger().Debugw("trace", "elapsed", elapsed, "rows", rows, "sql", sql)
	}
}

// nolint
var (
	gormPackage = filepath.Join("gorm.io")
	oilsPackage = filepath.Join("oils", "logs")
)

func (p *gormLogger) logger() *zap.SugaredLogger {
	start := 2
	end := 15

	for skip := start; skip < end; skip++ {
		_, file, _, ok := runtime.Caller(skip)

		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, oilsPackage):
		default:
			// log.Debugw("skip", "skip", skip, "file", file)
			return logs.Desugar().WithOptions(zap.AddCallerSkip(skip - 1)).Sugar()
		}
	}

	return logs.Desugar().Sugar()
}
