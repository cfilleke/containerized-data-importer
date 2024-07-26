package cdicloner

import (
	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
)

// SetupMetrics register prometheus metrics
func SetupMetrics() error {
	return operatormetrics.RegisterMetrics(
		clonerMetrics,
	)
}
