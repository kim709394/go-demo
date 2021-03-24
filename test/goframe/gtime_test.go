package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"testing"
)

func TestGtime(t *testing.T) {
	//追加一天
	tomrrow := gtime.Now().AddDate(0, 0, 1)
	g.Dump(tomrrow)
}
