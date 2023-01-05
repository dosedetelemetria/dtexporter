package dtexporter

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
)

func (e *exp) createMetrics(mp metric.MeterProvider) error {
	var err error
	meter := mp.Meter("dtexporter")

	e.perRegionCounter, err = meter.SyncInt64().Counter(
		"dtexporter_per_region_counter",
		instrument.WithDescription("Number of spans we sent per region"),
		instrument.WithUnit(unit.Dimensionless),
	)
	if err != nil {
		return err
	}

	return nil
}
