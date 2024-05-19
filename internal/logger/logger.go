package logger

import "go.uber.org/zap"

var log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	log = logger.Sugar()
}

func Get() *zap.SugaredLogger {
	return log
}
