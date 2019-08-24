package _706_Design_HashMap

type MyHashMap struct {
	m map[int]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		m: make(map[int]int),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	ret, ok := this.m[key]
	if !ok {
		ret = -1
	}
	return ret
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	delete(this.m, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
