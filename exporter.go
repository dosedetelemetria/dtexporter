package dtexporter

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
)

type exp struct {
	perRegionCounter syncint64.Counter
}

func NewExporter() *exp {
	return &exp{}
}

func (e *exp) consumeTraces(ctx context.Context, ld ptrace.Traces) error {
	return nil
}
