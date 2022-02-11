package helloworld

import (
	cobraHandler "github.com/kazmerdome/godome/pkg/module/provider/handler/cobra"
	"github.com/spf13/cobra"
)

type helloworldCobraHandler struct {
	cobraHandler.CobraHandlerConfig
	helloworldService HelloworldService
}

func newHelloworldCobraHandler(c cobraHandler.CobraHandlerConfig, helloworldService HelloworldService) cobraHandler.CobraHandler {
	h := helloworldCobraHandler{
		helloworldService:  helloworldService,
		CobraHandlerConfig: c,
	}
	return &h
}

func (r *helloworldCobraHandler) AddSubcommand() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "Say Hello",
		Long:  `Say Hello World`,
		Run: func(cmd *cobra.Command, args []string) {
			r.GetLogger().Info(r.helloworldService.SayHello())
		},
	}
}
