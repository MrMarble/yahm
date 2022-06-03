
<div align="center">
<a href="https://gopherize.me">
<img src="assets/gopher.png" height="120" alt="gopher with moustache">
</a>

### YAHM
Yet Another Hook Manager

[![golangci-lint](https://github.com/MrMarble/yahm/actions/workflows/golangci.yml/badge.svg)](https://github.com/MrMarble/yahm/actions/workflows/golangci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrmarble/yahm)](https://goreportcard.com/report/github.com/mrmarble/yahm)
![Lines of code](https://img.shields.io/tokei/lines/github/mrmarble/yahm)
</div>

---

YAHM is a simple cli to manage your git hooks without third dependencies like python or node.

## Why?

Because I wanted to. I also don't like using [pre-commit](https://pre-commit.com/) because I don't have python on all the computers I use for coding and having to rely on third party repositories to run my hooks is a little weird, especially for running simple commands like `go test ./...` or `npm test` (although husky is :ok_hand: if you use node).

I like my tools to be a single binary without system dependencies to use them anywhere without too much overhead.

## Usage

Just run `yahm install` on a folder containing the `.yahm.yaml` file (or pass the file using `-c`) and thats it

### Config file

This is the basic structure of the config file, check this repo [.yahm.yaml](.yahm.yaml) file for more examples:

```yaml
hooks:
  pre-commit:
    actions:
      - name: format
        cmd: gofumpt -w -extra ./..
      - name: test
        cmd: go test ./...
```

### Planned features
Things I'm interested in doing but may never implement lol

- Load from file
- Load from url
- Improve git output
- Built-in hooks. Like conventional commit messages
- Show installed/modified hook on `yahm list`
- Run in parallel?
