package main

import (
	"OtelMetricsQuickstartExample/opentelemetry"
	"time"
)

func main() {

	defer opentelemetry.Shutdown()
	opentelemetry.InitMetrics()

	for {
		time.Sleep(time.Second * 1)
		opentelemetry.SetHistogramValue("testVal", 1)
	}

}
