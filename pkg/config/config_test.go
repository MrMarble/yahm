package config_test

import (
	"testing"

	"github.com/mrmarble/ghooks/pkg/config"
	"github.com/stretchr/testify/assert"
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
	assert.NotNil(t, c)
	assert.NotNil(t, c.Hooks)
	assert.Len(t, c.Hooks, 2)
	assert.NotNil(t, c.Hooks["pre-commit"])
	assert.NotNil(t, c.Hooks["commit-msg"])
	assert.Len(t, c.Hooks["pre-commit"].Actions, 1)
	assert.Len(t, c.Hooks["commit-msg"].Actions, 1)
}
