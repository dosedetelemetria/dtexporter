package dtexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	typeStr   component.Type = "dtexporter"
	stability                = component.StabilityLevelDevelopment
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithTraces(createTraces, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Region: "ce",
	}
}

func createTraces(ctx context.Context, cs exporter.CreateSettings, cfg component.Config) (exporter.Traces, error) {
	return exporterhelper.NewTracesExporter(ctx, cs, cfg, consumeTraces)
}
