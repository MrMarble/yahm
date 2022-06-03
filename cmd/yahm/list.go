package main

import (
	"fmt"

	"github.com/mrmarble/yahm/pkg/config"
)

type listCmd struct{}

func (l *listCmd) Run(ctx *Context) error {
	cfg, err := config.Read(ctx.ConfigPath)
	if err != nil {
		return err
	}

	for kind, hook := range cfg.Hooks {
		for _, action := range hook.Actions {
			fmt.Printf("[%s] %s: %s\n", kind, action.Name, action.Cmd)
		}
	}

	return nil
}
