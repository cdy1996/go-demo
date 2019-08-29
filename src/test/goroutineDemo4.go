// 这个示例程序展示如何在程序里造成竞争状态
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter2 是所有goroutine 都要增加其值的变量
	counter2 int64

	// wg2 用来等待程序结束
	wg2 sync.WaitGroup

	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
)

// main 是所有Go 程序的入口
func main() {
	// 计数加 2，表示要等待两个goroutine
	wg2.Add(2)

	// 创建两个goroutine
	go inccounter2(1)
	go inccounter2(2)

	// 等待 goroutine 结束
	wg2.Wait()
	fmt.Println("Final counter2:", counter2)
}

// inccounter2 增加包里counter2 变量的值
func inccounter2(id int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			// 捕获 counter2 的值
			value := counter2

			// 当前 goroutine 从线程退出，并放回到队列
			runtime.Gosched()

			// 增加本地value 变量的值
			value++

			// 将该值保存回counter2
			counter2 = value
		}
		mutex.Unlock()

	}

}
