package monitors_client

import (
	"github.com/getsentry/sentry-go"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type MonitoringClient struct {
	Relic  *newrelic.Application
	Sentry *sentry.Client
}

func NewMonitoringClient() (client *MonitoringClient, err error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(""),
		newrelic.ConfigLicense(""),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	client.Relic = app
	return nil, nil
}
