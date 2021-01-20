package goframe

import (
	"github.com/gogf/gf/frame/g"
	"github.com/kim709394/go-demo/goframe/pojo"
	"testing"
)

/*
@Author kim
@Description  数据转换
@date 2021-1-20 17:54
*/

//g.Var泛型使用
func TestVar(t *testing.T) {
	groups := make([]pojo.Group, 5)
	groups = append(groups, pojo.Group{
		Id: 1,
	}, pojo.Group{
		Id: 2,
	})
	newVar := g.NewVar(groups)
	var i interface{} = newVar
	g.Dump(i)
}
