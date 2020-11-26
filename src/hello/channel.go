package hello

import (
	"fmt"
	"time"
)

/*
@Author kim
@Description  通道定义
@date 2020-11-25 12:52
*/

/*通道,用于协程之间的通信，类似于消息队列里面的队列，不同协程之间可以通过发送和接收通道的消息进行通信，通道发送方和接收方
*会达到平衡，否则如果发送方发送的数据永远没有接收方接收将会报异常，接收方会一直阻塞在通道直到能接收到数据
 */

//以下是无缓冲通道，要求发送方和接收方达到同步(平衡)
var cn chan (int) = make(chan int)

//发送和接收数据
func SendAndReceive() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		//往通道发送数据
		cn <- 110
	}()

	go func() {
		defer wg.Done()
		//阻塞接收通道的数据,
		data := <-cn
		//非阻塞接收通道数据,ok表示是否接收到数据，这种方式不会造成阻塞，但是会造成cpu过高，一般不使用
		//data,ok:=<-cn
		//忽略接收到的数据,只接收通道里的数据
		//<-cn
		fmt.Println(data)
	}()
	wg.Wait()
}

//循环接收通道数据,发送和接收必须同时匹配，只发送不接收或者只接收不发送都会被go语言智能发现，会死锁异常
func LoopRecv() {

	go func() {
		for i := 0; i < 10; i++ {
			cn <- i
		}
	}()

	for data := range cn {

		fmt.Println("接收到的数据", data)
		if data == 9 {
			break
		}
	}
}

//单向通道
func OneWayCn() {

	/*//创建只写单向通道
	write:=make(chan<- int)
	//创建只读单向通道
	read:=make(<-chan int)
	//往只写通道写入数据
	write<-100
	//往只读通道读取数据
	data:=<-read
	fmt.Println(data)
	单向通道其实没有意义，只是作为一种对通道在写和读的场合的约束
	*/
	cn := make(chan int)
	go func(write chan<- int) {
		for i := 0; i < 10; i++ {
			write <- i
		}
		close(write)
	}(cn)
	func(read <-chan int) {
		for {
			data, close := <-read
			if !close {
				fmt.Println("通道读取完毕")
				break
			}
			fmt.Println(data)
		}
	}(cn)

}

//关闭通道
func CloseCn() {

	cn := make(chan int)
	close(cn)
	//检查通道是否关闭
	_, close := <-cn
	if !close {
		fmt.Println("通道已关闭")
	}
}

//带缓冲通道，不强制发送方和接收方要同步，发送方发送的消息超过缓冲容量时会发生阻塞，接收方在缓冲通道的数据为空时会阻塞

func ChanWithCache() {
	//创建带缓冲通道，3为缓冲通道容量
	cnCh := make(chan int, 3)
	//缓冲容量大小
	fmt.Println("发消息前缓冲容量大小：", len(cnCh))
	for i := 0; i < 3; i++ {
		cnCh <- i
	}
	//缓冲容量大小
	fmt.Println("发送消息后缓冲容量大小：", len(cnCh))
	for i := 0; i < 3; i++ {
		data := <-cnCh
		fmt.Println("缓冲通道里的数据", data)
	}
	//缓冲容量大小
	fmt.Println("接收消息后缓冲容量大小:", len(cnCh))
}

//select语句
func Sele() {
	//超时函数
	after := time.After(3 * time.Second)
	data := <-after
	fmt.Println(data)
	/*select语句,与switch非常相似，多了一些限制，case语句必须是IO操作，如果找不到case匹配项，有default走default，
	*没有default则会阻塞住，直到有满足的case操作才会结束
	 */
	channel := make(chan int)
	timer := make(chan bool)
	go func() {
		select {
		case num := <-channel:
			fmt.Println(num)
		case channel <- 12:
			fmt.Println("写入数据")
		case <-time.After(3 * time.Second):
			fmt.Println("超时")
			timer <- true
		}
	}()
	<-timer
}
