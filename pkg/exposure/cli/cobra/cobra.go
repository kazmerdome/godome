package cobraExposure

import (
	"github.com/kazmerdome/godome/pkg/exposure"
	"github.com/spf13/cobra"
)

type CobraExposure interface {
	exposure.ExposureConfig
	Execute()
}

type Handler interface {
	AddSubcommand() *cobra.Command
}

type cobraExposure struct {
	exposure.ExposureConfig
	rootCmd *cobra.Command
}

func NewCobraExposure(
	c exposure.ExposureConfig,
	handlers []Handler,
	RootUse string,
	RootShort string,
	RootLong string,
) CobraExposure {
	e := new(cobraExposure)
	e.rootCmd = &cobra.Command{
		Use:   RootUse,
		Short: RootShort,
		Long:  RootLong,
	}
	for _, handler := range handlers {
		e.rootCmd.AddCommand(handler.AddSubcommand())
	}
	return e
}

func (r *cobraExposure) Execute() {
	r.rootCmd.Execute()
}
