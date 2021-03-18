package utils

func pop(alist *[]string) string {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = append((*alist)[:f-1])
	return rv
}
