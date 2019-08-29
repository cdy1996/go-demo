// 这个示例程序展示如何在程序里造成竞争状态
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是所有goroutine 都要增加其值的变量
	counter int64

	// wg1 用来等待程序结束
	wg1 sync.WaitGroup
)

// main 是所有Go 程序的入口
func main() {
	// 计数加 2，表示要等待两个goroutine
	wg1.Add(2)

	// 创建两个goroutine
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 结束
	wg1.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg1.Done()

	//for count := 0; count < 2; count++ {
	//	// 捕获 counter 的值
	//	value := counter
	//
	//	// 当前 goroutine 从线程退出，并放回到队列
	//	runtime.Gosched()
	//
	//	// 增加本地value 变量的值
	//	value++
	//
	//
	//	// 将该值保存回counter
	//	counter = value
	//}
	for count := 0; count < 2; count++ {
		// 原子+1
		atomic.AddInt64(&counter, 1)

		// 当前 goroutine 从线程退出，并放回到队列
		runtime.Gosched()

	}
}
