package collector

import (
	"github.com/ustkit/golang/magent/internal/metric"
)

type Console Collector

func (c *Console) Send(metrics metric.Metrics) (err error) {
	c.Logger.Info(metrics)

	return nil
}
