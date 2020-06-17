package main

import (
	. "container/list"
	"fmt"
)

func main() {
	cache := Constructor1(2)

	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Printf("%v\n", cache.Get(1)) // 返回  1
	//cache.Put(3, 3)                  // 该操作会使得密钥 2 作废
	//fmt.Printf("%v\n", cache.Get(2)) // 返回 -1 (未找到)
	//cache.Put(4, 4)                  // 该操作会使得密钥 1 作废
	//fmt.Printf("%v\n", cache.Get(1)) // 返回 -1 (未找到)
	//fmt.Printf("%v\n", cache.Get(3)) // 返回  3
	//fmt.Printf("%v\n", cache.Get(4)) // 返回  4

	fmt.Printf("%v\n", cache.Get(2))
	cache.Put(2, 6)
	fmt.Printf("%v\n", cache.Get(1)) // 返回  1
	cache.Put(1, 5)                  // 该操作会使得密钥 2 作废
	cache.Put(1, 2)                  // 该操作会使得密钥 2 作废
	fmt.Printf("%v\n", cache.Get(1)) // 返回 -1 (未找到)
	fmt.Printf("%v\n", cache.Get(2)) // 返回 -1 (未找到)

}

type LRUCache struct {
	Map  map[int]int
	List *List
	Len  int
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{
		Map:  make(map[int]int),
		List: New(),
		Len:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	i, j := this.Map[key]
	if j {
		element := find(this.List, key)
		if element == nil {
			this.List.PushFront(key)
		} else {
			this.List.MoveToFront(element)
		}
		return i
	} else {
		return -1
	}
}

func find(list *List, key int) *Element {
	for item := list.Front(); nil != item; item = item.Next() {
		//fmt.Printf("%v %v", element.Value.(int), key)
		if item.Value.(int) == key {
			return item
		}
	}
	return nil
}

func (this *LRUCache) Put(key int, value int) {

	i := len(this.Map)
	_, j := this.Map[key]
	if i >= this.Len && !j {
		back := this.List.Back()
		this.List.Remove(back)
		delete(this.Map, back.Value.(int))
	}
	this.Map[key] = value
	element := find(this.List, key)
	if element == nil {
		this.List.PushFront(key)
	} else {
		this.List.MoveToFront(element)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
