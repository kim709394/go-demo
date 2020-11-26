package hello

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
@Author kim
@Description  并发编程
@date 2020-11-24 12:46
*/
//利用协程容器登记和退出登记所有的子协程
var wg sync.WaitGroup

/*创建一个goroutine(协程)，可以理解为线程内部的子线程，是比线程轻量很多的轻量级线程，占用内存大概在2KB，线程占用内存大概是2MB，创建和销毁的消耗资源
*成本要比线程低很多,类似java的创建一个线程。
 */
func process(i int) {
	defer wg.Done() //协程执行完毕，协程容器-1
	time.Sleep(1000)
	fmt.Println("在协程内执行:", i)
}

func Process() {
	//启动10个协程执行
	for i := 0; i < 10; i++ {
		//go关键字加一个函数调用，新建并启动一个协程
		go process(i)
		wg.Add(1) //往协程容器登记一个协程进去+1
	}
	//等待子协程们执行完毕，协程容器清零，主协程便执行结束
	wg.Wait()
	fmt.Println("主协程执行完毕")
}

//使用匿名函数创建协程
func ProcessByAnon() {

	//启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println("匿名函数闭包内的协程:", j)
		}(i)
	}
	wg.Wait()
	fmt.Println("主协程执行完毕")
}

//并发编程的安全性
var count int

//runtime.Gosched(),暂停当前协程，阻塞进队列，让cpu切换到另外一个协程执行,通常用到子协程代码里面
//go build -race 查询协程的竞争状态
func GoSched() {

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				tmp := count
				//暂停该协程，等待另外一个协程执行
				runtime.Gosched()
				tmp++
				count = tmp
			}
		}()
	}
	wg.Wait()
	fmt.Println("count=", count)
	fmt.Println("主协程执行完毕")

}

//线程安全函数atomic
var counter int64

func Atomic() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func LoadAndStore() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		atomic.StoreInt64(&counter, 1)
		fmt.Println(atomic.LoadInt64(&counter))
	}()
	go func() {
		defer wg.Done()
		atomic.StoreInt64(&counter, 2)
		fmt.Println(atomic.LoadInt64(&counter))
	}()

	wg.Wait()

}

var mutex sync.Mutex

func add() {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

//协程互斥锁
func RejectLock() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()
	fmt.Println(counter)
}

//查询cpu核数
func NumCPU() {
	numCPU := runtime.NumCPU()
	fmt.Println("逻辑CPU核数：", numCPU)
	/*类似java线程池设置核心池的线程数，设置协程调度器的处理数,设置了时候，后面代码执行的多协程将会以这个设置为基础
	*对分配cpu核数进行多协程的程序计算，多核每核处理单任务则并行，单核处理多任务则并发
	 */
	/*<1：不修改任何数值。
	=1：单核心执行。
	>1：多核并发执行。*/
	maxProcs := runtime.GOMAXPROCS(numCPU)
	fmt.Println("最大调度处理数：", maxProcs)
}

//读写互斥锁，读锁不会阻塞读操作，会阻塞写操作，写锁会阻塞读和写的操作
var rwMutex sync.RWMutex

func RWMutex() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(read int) {
			defer wg.Done()
			if read/2 > 0 {
				rwMutex.RLock()
				fmt.Println(counter)
				rwMutex.RUnlock()
			} else {
				rwMutex.Lock()
				counter++
				rwMutex.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(counter)

}
