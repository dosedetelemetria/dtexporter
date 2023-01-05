package dtexporter

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.uber.org/zap"
)

type exp struct {
	perRegionCounter syncint64.Counter
	region           attribute.KeyValue
	logger           *zap.Logger
}

func NewExporter(logger *zap.Logger, region string) *exp {
	regionAttr := attribute.String("region", region)
	return &exp{
		logger: logger,
		region: regionAttr,
	}
}

func (e *exp) consumeTraces(ctx context.Context, ld ptrace.Traces) error {
	for i := 0; i < ld.ResourceSpans().Len(); i++ {
		rs := ld.ResourceSpans().At(i)
		for sI := 0; sI < rs.ScopeSpans().Len(); sI++ {
			ss := rs.ScopeSpans().At(sI)
			for spanI := 0; spanI < ss.Spans().Len(); spanI++ {
				span := ss.Spans().At(spanI)
				e.logger.Debug("span received",
					zap.Stringer("parentID", span.ParentSpanID()),
					zap.Stringer("traceID", span.TraceID()),
				)
				e.perRegionCounter.Add(ctx, 1, e.region)
			}
		}
	}

	e.logger.Info("recebemos um pacote de trechos", zap.Int("count", ld.SpanCount()))
	return nil
}
