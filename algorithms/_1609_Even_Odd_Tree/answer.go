package _1609_Even_Odd_Tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	return isEvenOddTreeRunLine(root)
}

func isEvenOddTreeRunLine(root *TreeNode) bool {
	var (
		evenLines, oddLines [][]int
		nodeLine            = []*TreeNode{root}
		lineNum             = 0
	)
	// prepare data
	for len(nodeLine) != 0 {
		var (
			tmpNodeLine []*TreeNode
			tmpNumLine  []int
		)
		for _, node := range nodeLine {
			if node.Left != nil {
				tmpNodeLine = append(tmpNodeLine, node.Left)
			}
			if node.Right != nil {
				tmpNodeLine = append(tmpNodeLine, node.Right)
			}
			tmpNumLine = append(tmpNumLine, node.Val)
		}
		if lineNum%2 == 0 {
			evenLines = append(evenLines, tmpNumLine)
		} else {
			oddLines = append(oddLines, tmpNumLine)
		}

		lineNum++
		nodeLine = tmpNodeLine
	}
	// check data
	for _, line := range oddLines {
		if !isOddLine(line) {
			return false
		}
	}
	for _, line := range evenLines {
		if !isEvenLine(line) {
			return false
		}
	}
	return true
}

func isEvenLine(line []int) bool {
	var cur = math.MinInt
	for _, num := range line {
		if num%2 == 0 || num <= cur {
			return false
		}
		cur = num
	}
	return true
}

func isOddLine(line []int) bool {
	var cur = math.MaxInt
	for _, num := range line {
		if num%2 == 1 || num >= cur {
			return false
		}
		cur = num
	}
	return true
}
