package dtexporter

import "go.opentelemetry.io/collector/component"

var _ component.Config = (*Config)(nil)

type Config struct {
	DataCenter string `mapstructure:"datacenter"`
	Region     string `mapstructure:"region"`
}
