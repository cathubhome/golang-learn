package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
beego orm针对驼峰命名会自动帮你转化成下划线字段，如果你定义了Struct名字为UserInfo，
那么转化成底层实现的时候是user_info，字段命名也遵循该规则
*/

//Model Struct
type Userinfo struct {
	Uid        int `orm:"column(uid);pk"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid     int      `orm:"column(uid);pk"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name    string   `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	// 设置默认数据库:数据库别名、数据库驱动、数据源、数据库的最大空闲连接、数据库的最大数据库连接（go version1.12以上）
	orm.RegisterDataBase("default", "mysql", "root:cathome@tcp(192.168.17.8:3306)/test?charset=utf8", 30, 30)
	//支持打印调试
	orm.Debug = true

	// 注册定义的 model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Post), new(Tag))

	// 创建 table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	//var userinfo Userinfo
	//userinfo.Username="golang"
	//userinfo.Departname ="programming"
	//userinfo.Created = time.Now()

	//插入数据，同时插入多个对象insertMulti
	//i, e := o.Insert(&userinfo)
	//fmt.Println("id:",i)
	//check(e)

	//查询数据start
	//example 1
	//var I Userinfo
	//I.Uid=1
	//read := o.Read(&I)
	//if read == orm.ErrNoRows{
	//	fmt.Println("can not find")
	//}else if read == orm.ErrMissPK{
	//	fmt.Println("cann not find pk")
	//}else{
	//	fmt.Println(I.Username, I.Departname, I.Created)
	//}

	//example 2
	var userInfo2 Userinfo
	qs := o.QueryTable(userInfo2) //返回QuerySeter
	//SELECT T0.`uid`, T0.`username`, T0.`departname` FROM `userinfo` T0 WHERE T0.`username` = ? LIMIT 1
	error := qs.Filter("username", "golang").One(&userInfo2, "uid", "username", "departname")
	check(error)

	//复杂查询参考：https://blog.csdn.net/yang731227/article/details/82503059

	//查询数据end

	if num, err := o.Delete(&User{Uid: 1}); err == nil {
		fmt.Println(num)
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
