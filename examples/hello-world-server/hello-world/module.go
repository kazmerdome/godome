package helloworld

import (
	"github.com/kazmerdome/godome/pkg/module"
	echoHandler "github.com/kazmerdome/godome/pkg/module/provider/handler/echo"
)

type HelloworldModule interface {
	GetService() HelloworldService
	GetEchoHandler() echoHandler.EchoHandler
}

type helloworldModule struct {
	moduleConfig          module.ModuleConfig
	helloworldService     HelloworldService
	helloworldEchoHandler echoHandler.EchoHandler
}

func NewHelloworldModule(moduleConfig module.ModuleConfig) HelloworldModule {
	m := new(helloworldModule)
	m.moduleConfig = moduleConfig
	providerConfig := m.moduleConfig.GetProviderConfig()

	// Service
	m.helloworldService = newUserService(providerConfig)

	// Handlers
	m.helloworldEchoHandler = newHelloworldEchoHandler(providerConfig, m.helloworldService)

	return m
}

func (r *helloworldModule) GetService() HelloworldService {
	return r.helloworldService
}

func (r *helloworldModule) GetEchoHandler() echoHandler.EchoHandler {
	return r.helloworldEchoHandler
}
