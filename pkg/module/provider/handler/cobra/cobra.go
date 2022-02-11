package cobraHandler

import (
	"github.com/kazmerdome/godome/pkg/module/provider"
	"github.com/spf13/cobra"
)

type CobraHandler interface {
	AddSubcommand() *cobra.Command
}

type CobraHandlerConfig interface {
	provider.ProviderConfig
}

type cobraHandlerConfig struct {
	provider.ProviderConfig
}

func NewCobraHandlerConfig(config provider.ProviderConfig) CobraHandlerConfig {
	return &cobraHandlerConfig{ProviderConfig: config}
}
