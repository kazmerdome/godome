package helloworld_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kazmerdome/godome/pkg/config"
	"github.com/kazmerdome/godome/pkg/module"
	cobraHandler "github.com/kazmerdome/godome/pkg/module/provider/handler/cobra"
	standardLogger "github.com/kazmerdome/godome/pkg/observer/logger/standard"

	helloworld "github.com/kazmerdome/godome/examples/hello-world-cli/hello-world"
)

type fixture struct {
	service      helloworld.HelloworldService
	cobraHandler cobraHandler.CobraHandler
}

func newFixture() *fixture {
	f := new(fixture)
	// setup
	c := config.NewConfig(config.MODE_GLOBALENV)
	l := standardLogger.NewStandardLogger()
	// init module
	module := helloworld.NewHelloworldModule(module.NewModuleConfig(l, c))
	f.service = module.GetService()
	f.cobraHandler = module.GetCobraHandler()

	return f
}

func TestHelloworldModule(t *testing.T) {
	assert := assert.New(t)

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	cmd := newFixture().cobraHandler.AddSubcommand()
	cmd.SetArgs([]string{"hello"})
	cmd.Execute()

	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	assert.Equal(out, "Helloworld Service has been called.\nHello world\n")
}
