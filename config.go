package dtexporter

import (
	"errors"

	"go.opentelemetry.io/collector/component"
)

var _ component.Config = (*Config)(nil)

// Config represents the interface with our users
type Config struct {
	DataCenter string `mapstructure:"datacenter"`
	Region     string `mapstructure:"region"`
}

func (c *Config) Validate() error {
	if c.DataCenter == "" {
		return errors.New("datacenter must be provided")
	}

	return nil
}
