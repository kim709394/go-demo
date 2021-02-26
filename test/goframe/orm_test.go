package goframe

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/kim709394/go-demo/goframe/pojo"
	"testing"
)

/*
@Author kim
@Description  数据库操作
@date 2021-1-8 13:47
*/

/*
数据库配置：
[database]
	[[database.分组名称]]
		host                 = "地址"
		port                 = "端口"
		user                 = "账号"
		pass                 = "密码"
		name                 = "数据库名称"
		type                 = "数据库类型(mysql/pgsql/mssql/sqlite/oracle)"
		role                 = "(可选)数据库主从角色(master/slave)，不使用应用层的主从机制请均设置为master"
		debug                = "(可选)开启调试模式"
		prefix               = "(可选)表名前缀"
		dryRun               = "(可选)ORM空跑(只读不写)"
		charset              = "(可选)数据库编码(如: utf8/gbk/gb2312)，一般设置为utf8"
		weight               = "(可选)负载均衡权重，用于负载均衡控制，不使用应用层的负载均衡机制请置空"
		linkinfo             = "(可选)自定义数据库链接信息，当该字段被设置值时，以上链接字段(Host,Port,User,Pass,Name)将失效，但是type必须有值"
		maxIdle              = "(可选)连接池最大闲置的连接数"
		maxOpen              = "(可选)连接池最大打开的连接数"
		maxLifetime          = "(可选，单位秒)连接对象可重复使用的时间长度"
		createdAt            = "(可选)自动创建时间字段名称"
		updatedAt            = "(可选)自动更新时间字段名称"
		deletedAt            = "(可选)软删除时间字段名称"
		timeMaintainDisabled = "(可选)是否完全关闭时间更新特性，true时CreatedAt/UpdatedAt/DeletedAt都将失效"
简化配置：
[database]
    [[database.default]]
        type = "mysql"
        link = "root:12345678@tcp(127.0.0.1:3306)/test"
    [[database.user]]
        type = "mysql"
        link = "mysql:root:12345678@tcp(127.0.0.1:3306)/user"

也可以简化为：

[database]
    [[database.default]]
        link = "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
    [[database.user]]
        link = "mysql:root:12345678@tcp(127.0.0.1:3306)/user"
注意以上每一项分组配置均可以是多个节点，支持负载均衡权重策略。如果不使用多节点负载均衡特性，仅使用配置分组特性，也可以简化为如下格式：

[database]
    [database.default]
        link = "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
    [database.user]
        link = "mysql:root:12345678@tcp(127.0.0.1:3306)/user"
如果仅仅是单数据库节点，不使用配置分组特性，那么也可以简化为如下格式：

[database]
    link = "mysql:root:12345678@tcp(127.0.0.1:3306)/test"

不同数据类型对应的link如下:
mysql	mysql: 账号:密码@tcp(地址:端口)/数据库名称	mysql
pgsql	pgsql: user=账号 password=密码 host=地址 port=端口 dbname=数据库名称	pq
mssql	mssql: user id=账号;password=密码;server=地址;port=端口;database=数据库名称;encrypt=disable	go-mssqldb
sqlite	sqlite: 文件绝对路径 (如: /var/lib/db.sqlite3)	go-sqlite3
oracle	oracle: 账号/密码@地址:端口/数据库名称	go-oci8

日志输出配置
gdb支持日志输出，内部使用的是glog.Logger对象实现日志管理，并且可以通过配置文件对日志对象进行配置。默认情况下gdb关闭了DEBUG日志输出，如果需要打开DEBUG信息需要将数据库的debug参数设置为true。以下是为一个配置文件示例：

[database]
    [database.logger]
        path   = "/var/log/gf-app/sql"
        level  = "all"
        stdout = true
    [database.primary]
        link   = "mysql:root:12345678@tcp(127.0.0.1:3306)/user_center"
        debug  = true
其中database.logger即为gdb的日志配置，当该配置不存在时，将会使用日志组件的默认配置，

需要注意哦：由于ORM底层都是采用安全的预处理执行方式，提交到底层的SQL与参数其实是分开的，因此日志中记录的完整SQL仅作参考方便人工阅读，并不是真正提交到底层的SQL语句。
*/
/*
mysql数据类型与go变量类型映射
mysql    	go变量
*char		string
*text		string
*binary		bytes
*blob		bytes
*int		int
*money		float64
bit			int
big_int		int64
float		float64
double		float64
decimal		float64
bool		bool
date		time.Time
datetime	time.Time
timestamp	time.Time
其他			string

数据库表字段名与go变量名之间的映射匹配规则：
表键名     字段名称     	是否匹配
nickname   nickname      match
NICKNAME   nickname      match
Nick-Name  nickname      match
nick_name  nickname      match
nick name  nickname      match
NickName   nickname      match
Nick-name  nickname      match
nick_name  nickname      match
nick name  nickname      match
*/

//获取数据库操作对象
func TestDBModel(t *testing.T) {
	//select * from t_group,t_phone
	//model := g.DB().Model("t_group,t_phone")
	model := g.DB().Table("t_phone") //等价于g.DB().Model("t_phone")
	res, err := model.Where("id in (?)", g.Slice{1, 2}).All()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	//model默认是链式不安全的,连续调用where则是并列关系，where id in (1,2) and create_time >= now()
	model.Where("created_at >= ?", gtime.Now()).All()
	//clone()方法创建一个新的model对象出来，保留原model对象的所有特性，与原model对象不关联
	clone := model.Clone()
	fmt.Println(clone)

	//safe()方法创建链式安全对象
	model1 := g.DB().Model("t_phone").Safe()
	//model2:where name = 'hello'
	model2 := model1.Where("name = ?", "hello")
	//model3:where name = 'hi'
	model3 := model1.Where("name = ?", "hi")
	//model2与model3两者非并列关系，两者独立不关联
	fmt.Println(model2, model3)
}

//实体，字段名和数据库的键名默认是驼峰匹配下划线规则,首字母大写是为了公开访问权限，不影响与数据库键名的首字母小写匹配
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

//插入操作
func TestInsert(t *testing.T) {
	/*Insert()方法：使用INSERT INTO语句进行数据库写入，如果写入的数据中存在主键或者唯一索引时，返回失败，否则写入一条新数据；
	InsertIgnore()方法：使用INSERT IGNORE INTO语句进行数据库写入，如果写入的数据中存在主键或者唯一索引时，忽略错误继续执行写入；
	Replace()方法：使用REPLACE INTO语句进行数据库写入，如果写入的数据中存在主键或者唯一索引时，会删除原有的记录，必定会写入一条新记录；
	Save()方法：使用INSERT INTO语句进行数据库写入，如果写入的数据中存在主键或者唯一索引时，更新原有数据，否则写入一条新数据；*/
	//使用Data()方法，参数为Map
	res, err := g.DB().Model("t_phone").Data(g.Map{"ip_address": "127.0.0.1", "name": "话机1", "phone_num": "110",
		"password": "123456", "created_at": gtime.Now(), "deleted_at": 0, "g_id": 1}).Insert()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	//不使用Data()方法
	phone := new(Phone)
	phone.Name = "话机11"
	phone.Id = 1
	phone.CreatedAt = gtime.Now()
	phone.IpAddress = "localhost"
	phone.GId = 1
	phone.Password = "123456"
	phone.PhoneNum = "10000"
	g.DB().Model("t_phone").Save(phone)
	//批量写入，insert into t_phone name values ('话机2'),('话机3')
	g.DB().Model("t_phone").Insert(g.List{{"name": "话机2"}, {"name": "话机3"}})
	//批量多条语句插入,Batch()方法，默认是10条语句
	//insert into t_phone name values ('话机2'),('话机3')
	//insert into t_phone name values ('话机4')
	g.DB().Model("t_phone").Batch(2).Insert(g.List{{"name": "话机2"}, {"name": "话机3"}, {"name": "话机4"}})

	//RawSQL嵌入,生成的sql语句不会转化成字符串，只会原班不动执行：gdb.Raw("now()")   now()不会变成'now()'
	g.DB().Model("t_phone").Data(g.Map{"ip_address": "127.0.0.1", "name": "话机1", "phone_num": "110",
		"password": "123456", "created_at": gdb.Raw("now()"), "deleted_at": 0, "g_id": 1}).Insert()
}

//修改方法
func TestUpdate(t *testing.T) {
	//个性化修改
	//g.DB().Model("t_phone").Data(g.Map{"name": "709394"}).Where("id=?", 1).Update()
	phone := new(Phone)
	phone.Name = "话机11111"
	phone.Id = 2
	phone.CreatedAt = gtime.Now()
	phone.IpAddress = "localhost"
	phone.GId = 1
	phone.Password = "123456"
	phone.PhoneNum = "10000"
	//主键存在的情况下进行修改，如果数据库字段约束设置了not null的话，空值无法忽略，此时结构体的每个属性必须有值或者该字段必须有默认值
	g.DB().Model("t_phone").OmitEmpty().Save(phone)

	g.DB().Model("t_group").Insert(g.Map{"name": "井下", "deleted_at": gtime.Now(), "created_at": gdb.Raw("now()")})
	//Counter对象用法，用于给字段修改为自增
	counter := &gdb.Counter{
		Field: "del",
		Value: 1,
	}
	// update t_group set del=del+1 where id=1
	g.DB().Model("t_group").Data(g.Map{"del": counter}).Where("id=?", 1).Update()
}

//删除
func TestDel(t *testing.T) {
	//delete from t_phone where id=5
	g.DB().Model("t_phone").Delete(g.Map{"id": 5})

}

//查询
func TestQuery(t *testing.T) {
	//原生操作
	all, err := g.DB().Model("t_phone").
		Where("id=?", 1).
		And("name like ?", "%井%").
		Or("id =?", 0).
		Or("phone_num in (?)", g.Slice{2}).All()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(all)
	//对象操作,OmitEmpty():忽略空值
	//select * from t_group where id = 1 limit 1
	one, err1 := g.DB().Model("t_group").OmitEmpty().Where(Group{Id: 1}).One()
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(one)

	//wherePri()方法：查询条件仅有主键
	//select * from t_group where id=1
	g.DB().Model("t_group").WherePri(1)
	//select * from t_group where id in (1,2)
	g.DB().Model("t_group").WherePri(g.Slice{1, 2})

	//All(),返回所有符合条件的记录
	g.DB().Model("t_group").WherePri(1).All()
	//One(),返回第一条记录
	g.DB().Model("t_group").WherePri(1).One()
	//Array():返回某一个字段的记录，是一个数组,select id from t_group where id=1
	g.DB().Model("t_group").Array("id", "id=?", 1)
	//Value():返回一个字段值
	g.DB().Model("t_group").Value("id", "id=?", 1)
	g.DB().Model("t_group").Value("id", "id in (?)", g.Slice{1, 2})
	//Fields():需要查询的字段,select sum(id) from t_group where id=1
	g.DB().Model("t_group").Fields("sum(id)").WherePri(1).Value()
	//select distinct id,name from t_group where id=1
	g.DB().Model("t_group").Fields("distinct id,name").WherePri(1).All()
	//FieldsEx():排除查询的字段，select del,name from t_group where id=1
	g.DB().Model("t_group").FieldsEx("id", "created_at").WherePri(1).One()
	//count(),返回总记录数
	g.DB().Model("t_group").WherePri(1).Count()

	g.DB()
	//查询结果转化
	//Struct()查询结果转化为一个对象，这种方式是先初始化对象(先分配内存)，然后再将查询结果赋值给该对象
	group := new(Group)
	g.DB().Model("t_group").WherePri(1).Struct(group)
	//这种方式是先声明变量但还未分配内存，然后再将查询结果出来后实例化对象分配内存再封装给该对象，推荐使用这种方式
	group = (*Group)(nil)
	g.DB().Model("t_group").WherePri(1).Struct(group)
	//Structs()：返回一个对象数组
	groups := ([]*Group)(nil)
	g.DB().Model("t_group").Where("id between ? and ?", 0, 1).Structs(&groups)
	//Scan()方法,自动识别调用struct()还是structs()方法，如果返回的结果集是数组则调用structs()方法，如果是单条记录则返回struct()方法。
	g.DB().Model("t_group").Scan(&group)
	/*Find*支持主键条件的数据查询
	Find*方法包含：FindAll/FindOne/FineValue/FindCount/FindScan，
	这些方法与All/One/Value/Count/Scan方法的区别在于，当方法直接给定条件参数时，
	前者的效果与WherePri方法一致；而后者的效果与Where方法一致。也就是说Find*方法的条件参数支持智能主键识别特性。*/

	//group by
	g.DB().Model("t_phone").Fields("count(1) total,name").Group("name").All()
	//order by
	g.DB().Model("t_phone").Order("id desc,del asc").All()
	//having
	g.DB().Model("t_phone").Fields("count(1) total,phone_num").Group("phone_num").
		Having("phone_num > ?", 100).All()
}

//limit 分页
func TestLimitPaging(t *testing.T) {
	groups := ([]*Group)(nil)
	//var groups []*Group
	g.DB().Model("t_group").Limit(1, 2).Order("created_at desc").Structs(&groups)
	g.Dump(groups)

}

//事务处理
//Transaction方式处理事务,可以集中处理多个独立数据处理方法，保持事务一致性，当出现异常时则回滚，不出异常则提交事务
func TestTransaction(t *testing.T) {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		addGroup(nil)
		addPhone(tx)
		panic("报个错") //出现异常时将回滚
		//return nil
	})
	if err != nil {
		panic(err)
	}

}

//tx链式操作
func TestTxLink(t *testing.T) {
	tx, err := g.DB().Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = addPhone(tx)
	_, err = addGroup(tx)
}

//新增组
func addGroup(tx *gdb.TX) (id int64, err error) {
	var model *gdb.Model
	if tx == nil {
		model = g.DB().Model("t_group")
	} else {
		model = tx.Model("t_group")
	}
	group := &Group{
		Name: "一组",

		CreatedAt: gtime.Now(),
	}
	res, err := model.Insert(group)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return res.LastInsertId()
}

//新增电话机
func addPhone(tx *gdb.TX) (id int64, err error) {
	var model *gdb.Model
	if tx == nil {
		//事务对象为空，方法内独立事务管理
		model = g.DB().Model("t_phone")
	} else {
		//事务对象由外部传入，由外部管理
		model = tx.Model("t_phone")
	}
	phone := &Phone{
		Name:      "话机22",
		IpAddress: "localhost",
		PhoneNum:  "10086",
		Password:  "709394",
		CreatedAt: gtime.Now(),
		DeletedAt: gtime.Now(),
		GId:       1,
	}
	res, err := model.Insert(phone)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return res.LastInsertId()
}

//时间更新,针对结构体不生效，只针对map参数传入有效
/*
表中设置三个字段：
字段名            字段类型                        字段说明
Updated_at	  date, datetime, timestamp		   修改时间，每次数据被修改时自动更新时间
Created_at	  date, datetime, timestamp		   创建时间，数据插入时写入时间，只会写入一次
Deleted_at	  date, datetime, timestamp		   删除时间，数据被删除时写入时间，实则为逻辑删除，查询的时候将会自带Deleted_at is not null语句
*/
func TestTimeChange(t *testing.T) {
	//针对结构体不生效
	res, err := g.DB().Model("t_group").Insert(Group{
		Name: "二组",
	})
	if err != nil {
		fmt.Println(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	g.DB().Model("t_group").Save(Group{Id: id, Name: "三组"})
	g.DB().Model("t_group").Delete("id =?", id)
}

//测试时间更新，参数为map，时间更新有效
func TestTimeChangeWithMap(t *testing.T) {

	res, err := g.DB().Model("t_group").Insert(g.Map{"name": "四组"})
	if err != nil {
		glog.Error(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		glog.Error(err)
	}
	g.DB().Model("t_group").Save(g.Map{"name": "五组", "id": id})
	g.DB().Model("t_group").Delete("id=?", id)

}

//查询结果处理
func TestResult(t *testing.T) {

	one, err := g.DB().Model("t_group").WherePri(1).One()
	fmt.Println(one)
	all, err := g.DB().Model("t_group").Where("id in (?)", g.Slice{1, 2, 3, 4, 5}).All()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range all {
		fmt.Println(v.GMap())
	}

}

//查询结果为空的测试用例
func TestNullResult(t *testing.T) {

	array, err := g.DB().Model("t_group").Where("id > 30").Fields("id").Array()
	//查询为空不报错
	if err != nil {
		g.Dump(err.Error())
	}
	//查询无记录则数组为空
	if array == nil || len(array) == 0 {
		g.Dump(array)
	}
	s := gconv.Int64s(array)
	g.Dump("输出结果:", s)
}

//查询记录为空
func TestNullRecord(t *testing.T) {
	group := new(pojo.Group)
	err := g.DB().Model("t_group").Where("id > 30").Struct(group)
	//查询无记录则会报错
	g.Dump("sql: no rows in result set" == err.Error())
	if err != nil {
		g.Dump("报错了:", err.Error())
	}
	g.Dump(group)
}

//查询记录为空时得出的count值
func TestNullCount(t *testing.T) {
	count, err := g.DB().Model("t_group").Where("id > 30").Count()
	if err != nil {
		g.Dump("err : ", err.Error())
	}
	g.Dump(count)

}
