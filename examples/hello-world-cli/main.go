package main

import (
	"github.com/kazmerdome/godome/pkg/config"
	"github.com/kazmerdome/godome/pkg/exposure"
	cobraExposure "github.com/kazmerdome/godome/pkg/exposure/cli/cobra"
	"github.com/kazmerdome/godome/pkg/module"
	standardLogger "github.com/kazmerdome/godome/pkg/observer/logger/standard"

	helloworld "github.com/kazmerdome/godome/examples/hello-world-cli/hello-world"
)

const PORT = "9090"

func main() {
	// Load Config
	c := config.NewConfig(config.MODE_GLOBALENV)

	// Load Observers
	logger := standardLogger.NewStandardLogger()

	// Load Modules
	moduleConfig := module.NewModuleConfig(logger, c)
	helloworldModule := helloworld.NewHelloworldModule(moduleConfig)

	// Load Exposers
	cobraExposure := cobraExposure.NewCobraExposure(
		exposure.NewExposureConfig(logger, c),
		[]cobraExposure.Handler{
			helloworldModule.GetCobraHandler(),
		},
		"myappname",
		"short desc",
		"long desc",
	)
	cobraExposure.Execute()
}
