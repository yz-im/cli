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

type RunFunc func(basename string) error

func NewApp(name string, os ...Option) *App {
	a := &App{}
	for _, o := range os {
		o(a)
	}
	a.buildCommand()
	return a
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

func Default() *App {
	return &App{}
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	if a.runFunc != nil {
		return a.runFunc(a.name)
	}
	return nil
}

func (a *App) buildCommand() {
	cmd := &cobra.Command{}
	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}
	a.cmd = cmd
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
