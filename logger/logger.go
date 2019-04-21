package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	l, _ := zap.NewProduction()
	defer l.Sync() // flushes buffer, if any
	Log = l.Sugar()
}
