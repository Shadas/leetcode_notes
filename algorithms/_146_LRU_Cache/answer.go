package _146_LRU_Cache

import "sync"

type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}

type LRUCache struct {
	Mutex     sync.Mutex
	Capacity  int
	DummyHead *Node
	DummyTail *Node
	StoreMap  map[int]int
}

func Constructor(capacity int) LRUCache {
	dh := &Node{}
	dt := &Node{
		Next: dh,
		Pre:  dh,
	}
	dh.Next, dh.Pre = dt, dt
	lc := LRUCache{
		StoreMap:  make(map[int]int),
		Capacity:  capacity,
		DummyHead: dh,
		DummyTail: dt,
		Mutex:     sync.Mutex{},
	}
	return lc
}

func (this *LRUCache) Get(key int) int {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	if x, ok := this.StoreMap[key]; ok {
		// 调整顺序
		tmpNode := this.DummyHead.Next
		for {
			if tmpNode.Key == key {
				dummy := &Node{
					Key:   tmpNode.Key,
					Value: tmpNode.Value,
					Next:  this.DummyHead.Next,
					Pre:   this.DummyHead,
				}
				this.DummyHead.Next.Pre = dummy
				this.DummyHead.Next = dummy
				tmpNode.Pre.Next = tmpNode.Next
				tmpNode.Next.Pre = tmpNode.Pre
				tmpNode = nil
				break
			}
			tmpNode = tmpNode.Next
		}
		return x
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	if _, ok := this.StoreMap[key]; ok { // 如果已经存在
		// 更新链表
		tmp := this.DummyHead.Next
		for {
			if tmp.Key == key {
				tmp.Next.Pre = tmp.Pre
				tmp.Pre.Next = tmp.Next
				tmp = nil
				break
			}
			tmp = tmp.Next
		}
		this.addNode(key, value)
	} else {
		if len(this.StoreMap) < this.Capacity { // 直接添加，无需删除
			this.addNode(key, value)
		} else { // 添加后需要删除
			// 找到最后一个
			del := this.DummyTail.Pre
			del.Pre.Next = del.Next
			del.Next.Pre = del.Pre
			delete(this.StoreMap, del.Key)
			del = nil

			// 添加新的
			this.addNode(key, value)
		}
	}
}

func (this *LRUCache) addNode(key, value int) {
	newone := &Node{
		Key:   key,
		Value: value,
		Pre:   this.DummyHead,
		Next:  this.DummyHead.Next,
	}
	this.DummyHead.Next.Pre = newone
	this.DummyHead.Next = newone
	this.StoreMap[key] = value
}
