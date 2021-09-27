package _973_K_Closest_Points_to_Origin

import "container/heap"

func kClosest(points [][]int, k int) [][]int {
	return kClosestWithPQ(points, k)
}

// 前k个，用堆/优先队列即可
func kClosestWithPQ(points [][]int, k int) [][]int {
	h := PQ{}
	for _, point := range points {
		heap.Push(&h, point)
	}
	ret := make([][]int, k)
	for i := 0; i < k; i++ {
		ret[i] = heap.Pop(&h).([]int)
	}
	return ret
}

type Item struct {
	Point []int
	Val   int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

// x = []int, len=2
func (pq *PQ) Push(x interface{}) {
	point := x.([]int)
	item := &Item{Point: point, Val: point[0]*point[0] + point[1]*point[1]}
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item.Point
}
