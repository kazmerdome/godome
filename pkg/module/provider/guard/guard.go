package guard

import (
	"github.com/kazmerdome/godome/pkg/module/provider"
)

type GuardConfig interface {
	provider.ProviderConfig
}
type guardConfig struct {
	provider.ProviderConfig
}

func NewGuardConfig(c provider.ProviderConfig) GuardConfig {
	return &guardConfig{ProviderConfig: c}
}
