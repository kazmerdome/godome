package helloworld

import (
	"github.com/kazmerdome/godome/pkg/module"
	cobraHandler "github.com/kazmerdome/godome/pkg/module/provider/handler/cobra"
)

type HelloworldModule interface {
	GetService() HelloworldService
	GetCobraHandler() cobraHandler.CobraHandler
}

type helloworldModule struct {
	moduleConfig           module.ModuleConfig
	helloworldService      HelloworldService
	helloworldCobraHandler cobraHandler.CobraHandler
}

func NewHelloworldModule(moduleConfig module.ModuleConfig) HelloworldModule {
	m := new(helloworldModule)
	m.moduleConfig = moduleConfig
	providerConfig := m.moduleConfig.GetProviderConfig()

	// Service
	m.helloworldService = newUserService(providerConfig)

	// Handlers
	m.helloworldCobraHandler = newHelloworldCobraHandler(providerConfig, m.helloworldService)

	return m
}

func (r *helloworldModule) GetService() HelloworldService {
	return r.helloworldService
}

func (r *helloworldModule) GetCobraHandler() cobraHandler.CobraHandler {
	return r.helloworldCobraHandler
}
