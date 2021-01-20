package pojo

import "github.com/gogf/gf/os/gtime"

/*
@Author kim
@Description
@date 2021-1-20 17:57
*/

type Phone struct {
	Id        int64       //主键
	IpAddress string      //ip地址
	Name      string      //名字
	PhoneNum  string      //电话号码
	Password  string      //密码
	CreatedAt *gtime.Time //时间
	DeletedAt *gtime.Time //删除时间
	UpdatedAt *gtime.Time //修改时间
	GId       int64       //组id
}

type Group struct {
	Id        int64       //主键
	Name      string      //名字
	CreatedAt *gtime.Time //时间
	DeletedAt *gtime.Time //删除时间
	UpdatedAt *gtime.Time //修改时间
}
