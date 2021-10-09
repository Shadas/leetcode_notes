package _331_Verify_Preorder_Serialization_of_a_Binary_Tree

import "strings"

func isValidSerialization(preorder string) bool {
	var (
		eles  = strings.Split(preorder, ",")
		stack []string
	)
	for _, ele := range eles {
		stack = append(stack, ele)
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			stack = stack[0 : len(stack)-3]
			stack = append(stack, "#")
		}
	}
	return len(stack) == 1 && stack[0] == "#"
}
