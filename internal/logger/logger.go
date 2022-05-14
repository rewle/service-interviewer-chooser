package logger

import (
	"github.com/rewle/service-select-participants/internal/config"
	"go.uber.org/zap"
)

func Init(cfg *config.Config) *zap.SugaredLogger {
	log, _ := zap.NewProduction()
	defer log.Sync()
	return log.Sugar()
}
