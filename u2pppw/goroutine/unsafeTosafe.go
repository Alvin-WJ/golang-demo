package main


import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	//counter：multiple thread下的共享资源，全局变量，相当于static变量
	counter int64
	wg      sync.WaitGroup
)

func addCountNotSafe() {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//这里counter是共享资源，全局变量，相当于static变量，multiple thread下有同步问题
		counter++
		//当前goroutine从线程退出
		runtime.Gosched()
	}
}

func addCountAtomic() {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// 原子操作
		atomic.AddInt64(&counter,1)
		//当前goroutine从线程退出
		runtime.Gosched()
	}
}

func addCountLock() {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//加上锁，进入临界区域
		mutex.Lock()
		{
			counter++
			//当前goroutine从线程退出
			runtime.Gosched()
		}
		//离开临界区，释放互斥锁
		mutex.Unlock()
	}
}

func testNotSafe() {
	wg.Add(2)
	//下面多线程操作这个counter共享变量，出现并发问题
	go addCountNotSafe()
	go addCountNotSafe()
	wg.Wait()
	fmt.Printf("counter: %d\n",counter)
}