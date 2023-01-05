package dtexporter

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	// prepare
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	require.NotNil(t, cm)

	id := component.NewIDWithName("dtexporter", "2").String()
	sub, err := cm.Sub(id)
	require.NoError(t, err)

	// test
	cfg := &Config{}
	err = component.UnmarshalConfig(sub, cfg)
	require.NoError(t, err)

	// verify
	assert.Equal(t, "sp-1", cfg.DataCenter)
	assert.Equal(t, "sp", cfg.Region)
}
