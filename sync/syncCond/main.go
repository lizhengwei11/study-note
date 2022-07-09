package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.Cond 标准库使用
/*func main() {
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	//主协程先获取锁

	cond.L.Lock()
	fmt.Println("1>主协程获取到了锁")
	go func() {
		//协程一开始无法获取锁
		cond.L.Lock()
		defer cond.L.Unlock()
		fmt.Println("3>该协程获取到了锁")
		time.Sleep(3 * time.Second)
		//通过notify进行广播通知
		cond.Broadcast()
		fmt.Println("4>该协程执行完毕，即将释放defer中的解锁操作")
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("2>主协程依旧抢占着获得得锁")

	//先释放锁，直到收到了notify，又进行加锁
	cond.Wait()
	//释放锁
	cond.L.Unlock()
	fmt.Println("Done")

}*/

//实现多收一发场景
func main() {
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	for i := 1; i < 6; i++ {
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Printf("No.%d Groutine Receive\n", i)
		}(i)
	}

	time.Sleep(3 * time.Second)

	cond.Broadcast() //notify通知所有的锁
	//cond.Signal() //notify随机通知一个
	time.Sleep(time.Second * 5)
}
