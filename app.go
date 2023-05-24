package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type App struct {
	name    string
	cmd     *cobra.Command
	runFunc RunFunc
}

type RunFunc func(a *App)

func NewApp(name string, os ...Option) *App {
	a := &App{}
	for _, o := range os {
		o(a)
	}
	a.buildCommand()
	return a
}

func WithRunFunc(run RunFunc) RunFunc {
	return func(a *App) {
		a.runFunc = run
	}
}

func Default() *App {
	return &App{}
}

func (a *App) buildCommand() {
	a.cmd = &cobra.Command{}
}

func (a *App) Command() *cobra.Command {
	return a.cmd
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}
