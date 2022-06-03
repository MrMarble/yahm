package main

import (
	"fmt"
	"os"

	"github.com/mrmarble/yahm/pkg/config"
	"github.com/mrmarble/yahm/pkg/git"
	"github.com/mrmarble/yahm/pkg/yahm"
)

type installCmd struct{}

func (i *installCmd) Run(ctx *Context) error {
	cfg, err := config.Read(ctx.ConfigPath)
	if err != nil {
		return err
	}

	root, err := git.GetGitRoot()
	if err != nil {
		return err
	}
	for kind, hook := range cfg.Hooks {
		fmt.Println("Installing hook", kind)
		f, err := os.OpenFile(root+"/hooks/"+string(kind), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)
		if err != nil {
			return err
		}
		defer f.Close()
		err = yahm.WriteHook(f, hook)
		if err != nil {
			return err
		}
	}

	return nil
}
