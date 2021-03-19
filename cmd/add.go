/*
Copyright © 2021 MrMarble

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mrmarble/ghooks/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds a git hook",
	Long:    `Add (ghooks add) creates a new git hook file with the specified command.`,
	Example: "ghooks add pre-commit npm run lint",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requieres at least two arguments")
		}
		if utils.ValidateHook(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid hook specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		command := strings.Join(args[1:], " ")
		if err := add(args[0], command); err != nil {
			fmt.Print(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(hook, command string) error {
	if !utils.ValidateHook(hook) {
		return errors.New("Invalid hook")
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	gitRoot, err := utils.GetGitProjectRoot(cwd)
	if err != nil {
		return err
	}
	cmd := "#!/bin/sh" + "\n" + command
	hookPath := filepath.Join(gitRoot, "hooks", hook)
	return os.WriteFile(hookPath, []byte(cmd), os.ModePerm)
}
