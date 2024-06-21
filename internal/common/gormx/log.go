package gormx

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
	"time"
)

type Log struct {
	LogLevel logger.LogLevel
	Debug    bool
}

func NewLog(debug bool) *Log {
	return &Log{
		Debug: debug,
	}
}

func (l Log) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return &l
}

func (l Log) Info(ctx context.Context, s string, i ...interface{}) {
	fmt.Println("info", s, i)
}

func (l Log) Warn(ctx context.Context, s string, i ...interface{}) {
	fmt.Println("warn", s, i)
}

func (l Log) Error(ctx context.Context, s string, i ...interface{}) {
	fmt.Println("error", s, i)
}

func (l Log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.Debug {
		sql, rows := fc()
		logx.WithContext(ctx).Infof("cost: %vs rows: %v sql: %v", float64(time.Since(begin).Nanoseconds())/1e6, rows, sql)
	}
}
