package yahm

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

type HookType string

const (
	PreCommit HookType = "pre-commit"
	CommitMsg HookType = "commit-msg"
	PrePush   HookType = "pre-push"
	PreMergeCommit HookType = "pre-merge-commit"
	PrepareCommitMsg HookType = "prepare-commit-msg"
	// TODO: add more hooks
)

type Hook struct {
	Type    HookType
	Actions []action
}

type action struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

func NewHook(kind HookType) *Hook {
	return &Hook{
		Type:    kind,
		Actions: []action{},
	}
}

func (h *Hook) AddAction(name, cmd string) {
	h.Actions = append(h.Actions, action{
		Name: name,
		Cmd:  cmd,
	})
}

// WriteHook generates a git hook file
func WriteHook(w io.Writer, h *Hook) error {
	_, err := w.Write([]byte(fmt.Sprintf(`#!/usr/bin/env bash
# File generated by yahm: https://github.com/mrmarble/yahm
# At: %s

`, time.Now())))
	if err != nil {
		return err
	}
	for _, a := range h.Actions {
		_, err := w.Write([]byte(fmt.Sprintf(`
# action_start
# Name: %s
%s
# action_end
`, a.Name, a.Cmd)))
		if err != nil {
			return err
		}
	}
	return nil
}

func ParseHook(kind HookType, str string) (*Hook, error) {
	h := NewHook(kind)
	scanner := bufio.NewScanner(strings.NewReader(str))
	capturing := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# action_start") {
			h.Actions = append(h.Actions, action{})
			capturing = true
		} else if strings.HasPrefix(line, "# Name: ") {
			h.Actions[len(h.Actions)-1].Name = strings.TrimPrefix(line, "# Name: ")
		} else if strings.HasPrefix(line, "# action_end") {
			capturing = false
			h.Actions[len(h.Actions)-1].Cmd = strings.TrimSpace(h.Actions[len(h.Actions)-1].Cmd)
			continue
		} else if capturing {
			h.Actions[len(h.Actions)-1].Cmd += line + "\n"
		}
	}
	return h, nil
}
