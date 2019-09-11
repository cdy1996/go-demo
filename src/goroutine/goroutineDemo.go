package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//分配一个逻辑处理器给调用器使用
	runtime.GOMAXPROCS(1)
	//wg用于等待程序完成
	//计数+2, 表示有2个等待
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutine")

	// 声明定函数并创建goroutine
	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		//在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()
	fmt.Println("wait to finish")
	wg.Wait()
	fmt.Println("")
	fmt.Println("terminating")

}
