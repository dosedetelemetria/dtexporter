package dtexporter

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

func consumeTraces(ctx context.Context, ld ptrace.Traces) error {
	return nil
}
