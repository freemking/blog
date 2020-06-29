package main

import (
	"fmt"
	"time"
	"context"
)

var key string = "name"
var traceId string = "traceId"

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// //附加值
	// valueCtx := context.WithValue(ctx, key, "【监控1】")
	// go watch(valueCtx)
	// go watch(valueCtx)
	// go watch(valueCtx)
	// time.Sleep(10 * time.Second)
	// fmt.Println("可以了，通知监控停止")
	// cancel()
	// //为了检测监控过是否停止，如果没有监控输出，就表示停止了
	// time.Sleep(5 * time.Second)

	timeoutCtx,timeoutCancel := context.WithTimeout(context.Background(), time.Second)
	defer timeoutCancel()
	vCtx := context.WithValue(timeoutCtx, traceId, "123")
	go timeout(vCtx)
	time.Sleep(1100 * time.Millisecond)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func timeout(ctx context.Context) {
	ch := make(chan struct{}, 0)
    go func() {
        // 模拟4秒耗时任务
		time.Sleep(time.Second * 4)
		// 模拟0.1秒耗时任务
        // time.Sleep(time.Millisecond * 100)
        ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println(ctx.Value(traceId),"done")
	case <-ctx.Done():
		fmt.Println(ctx.Value(traceId),"timeout")
	}

}
