package _705_Design_HashSet

type MyHashSet struct {
	m map[int]interface{}
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		m: make(map[int]interface{}),
	}
}

func (this *MyHashSet) Add(key int) {
	this.m[key] = nil
}

func (this *MyHashSet) Remove(key int) {
	delete(this.m, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.m[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
