package collector

import (
	"github.com/ustkit/golang/magent/internal/config"
	"github.com/ustkit/golang/magent/internal/logger"
)

type Collector struct {
	Config *config.Common
	Logger *logger.Logger
}
