package commands

import (
	"github.com/Kdag-K/kdag/src/config"
)

// CLIConfig contains configuration for the Run command
type CLIConfig struct {
	Kdag     config.Config `mapstructure:",squash"`
	ProxyAddr  string        `mapstructure:"proxy-listen"`
	ClientAddr string        `mapstructure:"client-connect"`
}

// NewDefaultCLIConfig creates a CLIConfig with default values
func NewDefaultCLIConfig() *CLIConfig {
	return &CLIConfig{
		Kdag:     *config.NewDefaultConfig(),
		ProxyAddr:  "127.0.0.1:1338",
		ClientAddr: "127.0.0.1:1339",
	}
}
