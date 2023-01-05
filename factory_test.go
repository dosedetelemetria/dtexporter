package dtexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/exporter/exportertest"
)

func TestCreateDefaultConfig(t *testing.T) {
	// prepare
	expected := &Config{
		Region: "ce",
	}
	// test
	actual := createDefaultConfig()

	// verify
	assert.Equal(t, expected, actual)
}

func TestNewExporter(t *testing.T) {
	// prepare
	cfg := &Config{
		Region: "ce",
	}
	cs := exportertest.NewNopCreateSettings()
	f := NewFactory()

	// test
	exp, err := f.CreateTracesExporter(context.Background(), cs, cfg)
	require.NoError(t, err)

	// verify
	assert.NotNil(t, exp)
}
