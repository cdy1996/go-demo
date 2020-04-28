package main

import (
	"fmt"
	"golang.org/x/net/context"
	. "golang.org/x/time/rate"
	"sync"
	"time"
)

func main() {
	//每100ms往桶中放一个Token。本质上也就是一秒钟产生10个。
	limit := Every(100 * time.Millisecond)
	// 每秒token个数, 桶的大小
	limiter := NewLimiter(limit, 1)

	/*go func() {
		if limiter.AllowN(time.Now(), 10) {
			fmt.Printf("1 获取到了token")
		}
	}()
	go func() {
		if limiter.AllowN(time.Now(), 10) {
			fmt.Printf("2 获取到了token")
		}
	}()*/

	ctx, _ := context.WithCancel(context.Background())
	int1 := make(chan int, 1)
	int2 := make(chan int, 2)
	var wg sync.WaitGroup
	go func(ctx context.Context) {
		defer wg.Add(-1)
		wg.Add(1)
		int2 <- 1
		fmt.Printf("1 准备好了\n")
		<-int1
		fmt.Printf("1 准备开始\n")
		limiter.WaitN(ctx, 10)
		fmt.Printf("1 获取到了token\n")

	}(ctx)
	go func(ctx context.Context) {
		defer wg.Add(-1)
		wg.Add(1)
		int2 <- 2
		fmt.Printf("2 准备好了\n")
		<-int1
		fmt.Printf("2 准备开始\n")
		limiter.WaitN(ctx, 10)
		fmt.Printf("2 获取到了token\n")

	}(ctx)

	for i := 0; i != 2; {
		select {
		case n := <-int2:
			i++
			fmt.Printf("准备完毕 --- %d\n", n)

		}
	}
	int1 <- 1
	int1 <- 1
	wg.Wait()
	////调整token的速率
	//limiter.SetLimit(10)
	//
	////改变token 桶的大小
	//limiter.SetBurst(2)
}
