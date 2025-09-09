package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	LogLevel logger.LogLevel
}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		log.Printf("[GORM][INFO] %s", fmt.Sprintf(msg, data...))
	}
}

func (l *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		log.Printf("[GORM][WARN] %s", fmt.Sprintf(msg, data...))
	}
}

func (l *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		log.Printf("[GORM][ERROR] %s", fmt.Sprintf(msg, data...))
	}
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel == logger.Silent {
		return
	}

	sql, rows := fc()
	elapsed := time.Since(begin)

	if err != nil {
		log.Printf("[GORM][ERROR] %s [%.2fms] [rows:%v] %s",
			err, float64(elapsed.Microseconds())/1000.0, rows, sql)
	} else {
		log.Printf("[GORM][DEBUG] [%.2fms] [rows:%v] %s",
			float64(elapsed.Microseconds())/1000.0, rows, sql)
	}
}
