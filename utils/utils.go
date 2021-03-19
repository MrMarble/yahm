package utils

var hooks = map[string]struct{}{
	"applypatch-msg":     {},
	"commit-msg":         {},
	"fsmonitor-watchman": {},
	"post-update":        {},
	"pre-applypatch":     {},
	"pre-commit":         {},
	"pre-merge-commit":   {},
	"pre-push":           {},
	"pre-rebase":         {},
	"pre-receive":        {},
	"prepare-commit-msg": {},
	"update":             {},
}

func pop(alist *[]string) string {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func ValidateHook(hook string) bool {
	_, ok := hooks[hook]
	return ok
}
