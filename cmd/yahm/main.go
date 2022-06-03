package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

type VersionFlag string

type Context struct {
	Debug      bool
	ConfigPath string
}

var (
	// Populated by goreleaser during build
	version = "master"
	commit  = "?"
	date    = ""
)

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong) error {
	fmt.Printf("yahm has version %s built from %s on %s\n", version, commit, date)
	app.Exit(0)

	return nil
}

func main() {
	var cli struct {
		Debug   bool        `help:"Enable debug mode."`
		Config  string      `short:"c" type:"existingfile" help:"Path to config file." optional:""`
		Version VersionFlag `name:"version" help:"Print version information and quit"`

		Install installCmd `cmd:"install" help:"Install yahm to your git hooks directory"`
		List    listCmd    `cmd:"list" help:"List all hooks and their actions"`
	}

	ctx := kong.Parse(&cli,
		kong.Name("yahm"),
		kong.Description("A cli for managing git hooks"),
		kong.UsageOnError())

	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
