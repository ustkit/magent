package collector

import (
	"github.com/ustkit/magent/internal/config"
	"github.com/ustkit/magent/internal/logger"
)

type Collector struct {
	Config *config.Common
	Logger *logger.Logger
}
