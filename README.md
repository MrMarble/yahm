
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

YAHM is a simple cli to manage your git hooks without third dependencies.

## Why?

Because I wanted to. I also don't like using [pre-commit](https://pre-commit.com/) because I don't have python on all the computers I use for coding and having to rely on third party repositories to run my hooks is a little weird, especially for running simple commands like `go test ./...` or `npm test` (although husky is :ok_hand: if you use node).
