package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Profile))

	driverName := beego.AppConfig.String("dev::drivername")
	username := beego.AppConfig.String("dev::mysqluser")
	password := beego.AppConfig.String("dev::mysqlpass")
	dbUrl := beego.AppConfig.String("dev::mysqlurls")
	dbName := beego.AppConfig.String("dev::mysqldb")

	orm.RegisterDriver(driverName, orm.DRMySQL)

	//orm.RegisterDataBase("default", driverName, "root:root@/orm_test?charset=utf8")
	orm.RegisterDataBase("default", driverName, username+":"+password+"@"+dbUrl+"/"+dbName)
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // Using default, you can use other database

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
