package hello

import (
	"github.com/kim709394/go-demo/hello"
	"testing"
)

/*
@Author kim
@Description  并发编程测试
@date 2020-11-24 12:55
*/

//函数构建协程
func TestProcess(t *testing.T) {
	hello.Process()
}

//匿名函数构建协程
func TestProcessByAnon(t *testing.T) {
	hello.ProcessByAnon()
}

//协程调度和切换
func TestGoSched(t *testing.T) {
	hello.GoSched()
}

//原子性函数包atomic
func TestAtomic(t *testing.T) {
	hello.Atomic()
}

//线程安全原子性函数atomic.store,load
func TestLoadAndStore(t *testing.T) {
	hello.LoadAndStore()
}

//互斥锁
func TestMutex(t *testing.T) {
	hello.RejectLock()
}

//调度器最大处理数和CPU核心数
func TestCPUNum(t *testing.T) {
	hello.NumCPU()
}

//互斥读写锁
func TestRWMutex(t *testing.T) {
	hello.RWMutex()
}

//goexit
func TestGoexit(t *testing.T) {
	hello.GoExit()
}
