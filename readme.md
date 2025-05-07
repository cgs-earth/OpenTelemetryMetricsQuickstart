# OpenTelemetry Metrics Quickstart Example

This is a dummy quickstart example for OpenTelemetry Metrics. Since OpenTelemetry metrics requires more config to get up and running, this repository can help as a useful bootstrap.

It sends sends metrics from a golang program to the OpenTelemetry collector, which is then scraped by Prometheus and displayed in Grafana.

Run the example script with [run.sh](run.sh) and you can then create a dashboard in Grafana.