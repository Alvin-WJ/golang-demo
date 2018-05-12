package main

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

func test1() {
	for i := 0; i < 2; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}


func test2(){

	//分配一个逻辑处理器Ｐ给调度器使用（指定使用1个cpu的核来执行这段程序，以下定义的两个goroutine将被这一个cpu核来调度并发执行。一个CPU）
	runtime.GOMAXPROCS(1)
	//在这里,wg用于等待程序完成，计数器加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	//声明1个匿名函数，并创建一个goroutine
	fmt.Printf("Begin Coroutines\n")
	go func() {
		//在函数退出时，wg计数器减1
		defer wg.Done()
		//打印3次小写字母表
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			time.Sleep(10000)
		}
	}()
	//声明1个匿名函数，并创建一个goroutine
	go func() {
		defer wg.Done()
		//打印大写字母表3次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			time.Sleep(10000)
		}
	}()

	fmt.Printf("Waiting To Finish\n")
	//等待2个goroutine执行完毕
	wg.Wait()

}

func main() {
	//test1();
	test2()
}