package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//go func(ctx context.Context) {
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			fmt.Println("监控退出，停止了...")
	//			return
	//		default:
	//			fmt.Println("goroutine监控中...")
	//			time.Sleep(2 * time.Second)
	//		}
	//	}
	//}(ctx)
	//
	//time.Sleep(10 * time.Second)
	//fmt.Println("可以了，通知监控停止")
	//cancel()
	////为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)

	eatNum := chiHanBao(ctx)
	for n := range eatNum {
		if n >= 10 {
			cancel()
			break
		}
	}

	fmt.Println("正在统计结果。。。")
	time.Sleep(1 * time.Second)

}

func chiHanBao(ctx context.Context) <-chan int {
	c := make(chan int)
	// 个数
	n := 0
	// 时间
	t := 0
	go func() {
		for {
			//time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Printf("耗时 %d 秒，吃了 %d 个汉堡 \n", t, n)
				return
			case c <- n:
				incr := rand.Intn(5)
				n += incr
				if n >= 10 {
					n = 10
				}
				t++
				fmt.Printf("我吃了 %d 个汉堡\n", n)
			}
		}
	}()

	return c
}
