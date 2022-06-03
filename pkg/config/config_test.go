package config_test

import (
	"testing"

	"github.com/mrmarble/yahm/pkg/config"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalConfig(t *testing.T) {
	cfg := `hooks:
  pre-commit:
    actions:
      - name: lint
        cmd: gofmt -s -d .
  commit-msg:
    actions:
      - name: lint
        cmd: gofmt -s -d .`

	var c config.Config
	err := yaml.Unmarshal([]byte(cfg), &c)
	if err != nil {
		t.Fatal(err)
	}
	require.NotNil(t, c)
	require.NotNil(t, c.Hooks)
	require.Len(t, c.Hooks, 2)
	require.NotNil(t, c.Hooks["pre-commit"])
	require.NotNil(t, c.Hooks["commit-msg"])
	require.Len(t, c.Hooks["pre-commit"].Actions, 1)
	require.Len(t, c.Hooks["commit-msg"].Actions, 1)
}
