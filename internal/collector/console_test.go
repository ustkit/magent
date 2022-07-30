package collector

import (
	"testing"

	"github.com/ustkit/magent/internal/config"
	"github.com/ustkit/magent/internal/logger"
	"github.com/ustkit/magent/internal/metric"
)

func TestConsoleSend(t *testing.T) {
	metrics := make(metric.Metrics, 6)
	tlogger := logger.NewLogger(&config.Config{Common: config.Common{HostName: "localhost"}}, nil)
	console := &Console{Logger: tlogger}
	err := console.Send(metrics)
	if err != nil {
		t.Error(err)
	}
}
