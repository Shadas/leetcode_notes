package _71_Simplify_Path

import "strings"

func simplifyPath(path string) string {
	var (
		paths = strings.Split(path, "/")
		stack = []string{}
		ret   string
	)
	for _, path := range paths {
		if path == "" { // 空串往后走
			continue
		}
		if path == "." { // 本目录往后走
			continue
		}
		if path == ".." { // 向前一个目录
			if len(stack) > 0 { // 如果有元素，pop一个
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, path)
	}
	ret = strings.Join(stack, "/")
	ret = "/" + ret
	return ret
}
