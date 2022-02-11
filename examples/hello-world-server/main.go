package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kazmerdome/godome/pkg/config"
	"github.com/kazmerdome/godome/pkg/exposure"
	echoExposure "github.com/kazmerdome/godome/pkg/exposure/http/echo"
	"github.com/kazmerdome/godome/pkg/module"
	standardLogger "github.com/kazmerdome/godome/pkg/observer/logger/standard"

	helloworld "github.com/kazmerdome/godome/examples/hello-world-server/hello-world"
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
	echoExposure := echoExposure.NewEchoExposure(
		exposure.NewExposureConfig(logger, c),
		nil,
		[]echoExposure.Handler{
			helloworldModule.GetEchoHandler(),
		},
		PORT,
		true,
	)

	echoExposure.Start()
	defer echoExposure.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
