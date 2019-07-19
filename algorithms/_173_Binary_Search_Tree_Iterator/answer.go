package _173_Binary_Search_Tree_Iterator

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	l []int
}

func Constructor(root *TreeNode) BSTIterator {
	var (
		l    = []int{}
		s    = []*TreeNode{}
		node = root
	)
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		if len(s) > 0 {
			node = s[len(s)-1]
			s = s[:len(s)-1]
			l = append(l, node.Val)
			node = node.Right
		}
	}
	it := BSTIterator{
		l: l,
	}
	return it
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	var ret int
	if len(this.l) > 0 {
		ret = this.l[0]
		if len(this.l) > 1 {
			this.l = this.l[1:]
		} else {
			this.l = this.l[0:0]
		}

	} else {
		ret = 0
	}
	return ret
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.l) > 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
