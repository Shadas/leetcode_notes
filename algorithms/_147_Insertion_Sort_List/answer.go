package _147_Insertion_Sort_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next  // p为遍历剩下的链
	head.Next = nil // 拿出头部第一个节点
	for p != nil {
		pNext := p.Next     // 循环中，pNext 暂存后面剩下的链条, 此时 p直接为当前对比的节点，不需要考虑剩下的遍历
		q := head           // q 赋值为已经排序好的链
		if q.Val >= p.Val { // 如果待插入节点，比已排序节点的第一个节点还要小，则插入到最前面
			p.Next = q
			head = p
		} else { // 否则，应该向后找已排序节点的下一个节点进行比较
			for q != nil && q.Next != nil && q.Next.Val < p.Val { // 如果已排序链有下一个节点，并且下一个节点的值也小于待排序节点，则继续向后找
				q = q.Next
			}
			// 跳出循环，要么是遍历完了，要么是找到比待排序节点大的节点了，但总之当前的q后面，应该是插入p的位置
			p.Next = q.Next
			q.Next = p
		}
		// 插入完成，继续看下一个待排序节点
		p = pNext
	}
	return head
}
