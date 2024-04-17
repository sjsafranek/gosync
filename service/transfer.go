package service

import (
	"os"
	"time"
)

type transfer struct {
	File          *os.File
	StartTime     time.Time
	UpdateTime    time.Time
	BytesExpected int64
	BytesWritten  int64
}