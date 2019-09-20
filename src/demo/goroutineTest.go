package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			// 单独拥有副本
			fmt.Println(i)
		}(i)
		go func() {
			// 自由变量
			fmt.Println(i)
		}()
	}
	time.Sleep(500)

}

// 让协程 按顺序输出, 通过原子变量的递增
func testSyncGo() {
	count := uint32(0)

	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}

	trigger(10, func() {})
}
