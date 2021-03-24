package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"testing"
	"time"
)

//定时任务
func TestCronTimer(t *testing.T) {
	//cron := gcron.New()
	gcron.SetLogLevel(glog.LEVEL_ALL)
	_, err := gcron.AddOnce("*/5 * * * * *", func() {
		g.Dump("定时任务执行了：", gtime.Now())
	}, "interval1")
	if err != nil {
		g.Dump(err)
	}
	g.Dump("注册信息", gcron.Entries())
	//for{}
	time.Sleep(10 * time.Second)
	gcron.Remove("interval1")
}

// 2021-03-24 16:48:00
