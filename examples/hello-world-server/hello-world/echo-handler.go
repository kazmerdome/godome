package helloworld

import (
	"net/http"

	echoHandler "github.com/kazmerdome/godome/pkg/module/provider/handler/echo"
	"github.com/labstack/echo"
)

type helloworldEchoHandler struct {
	echoHandler.EchoHandlerConfig
	helloworldService HelloworldService
}

func newHelloworldEchoHandler(c echoHandler.EchoHandlerConfig, helloworldService HelloworldService) echoHandler.EchoHandler {
	h := helloworldEchoHandler{
		helloworldService: helloworldService,
		EchoHandlerConfig: c,
	}
	return &h
}

func (r *helloworldEchoHandler) AddSubroute(e *echo.Echo) {
	e.GET("/hello", func(c echo.Context) error {
		data := r.helloworldService.SayHello()
		return c.String(http.StatusOK, data)
	})
}
