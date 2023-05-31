package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// import "gorm.io/driver/mysql"

// 写一个user结构体
type User struct {
	gorm.Model //会帮忙建id 、create and update time
	Name       string
	Age        uint
}

func main() {

	/**
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

	*/

	// testMysql()

}

func testMysql() {
	//然后将mysql植进来,配置dsn属性
	//名称:密码@tcp(url:port)/
	dsn := "root:715705@Qc123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True"
	//然后用dsn和渠道，拿到db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//如果报错，
	if err != nil {
		panic("failed to connect database")
	}

	//迁移schema,如果没有建表，自动创建表
	db.AutoMigrate(&User{})
	//insert插入语句

	// ** Insert 插入语句
	// db.Create(&User{Name: "zhagnsan", Age: 20})

	// ** Select 查询语句
	var user User
	// db.First(&user, 1) // 根据主键查找
	// fmt.Println(user)
	//注意看第一个参数，这个是赋值
	db.First(&user, "name = ?", "linzy") // 查找 name 字段值为 linzy 的记录
	fmt.Println(user)

	var users []User
	//注意看第一个参数，这个是赋值操作
	result := db.Find(&users) // SELECT * FROM users;
	if result.Error != nil {
		// handle error
		fmt.Println("selct error")
	}
	// for _, user := range users {
	// 	fmt.Println(user) // print user
	// }
	fmt.Println(users)

	/**
	db.First(&user, 1) // 根据主键查找第一条记录1
	db.Take(&user) // 随机获取一条记录1
	db.Last(&user) // 根据主键查找最后一条记录1
	db.Find(&users) // 查找所有匹配的记录1
	db.Where(“name = ?”, “jinzhu”).First(&user) // 根据条件查找第一条记录1
	db.Where(“name = ?”, “jinzhu”).Find(&users) // 根据条件查找所有记录1
	*/

	// ** update 更新语句
	// Update 更新语句 - 将 User 的 age 更新为 18
	// db.Model(&user).Update("Age", 16) //更新了第一条
	db.Model(&user).Updates(User{Name: "linzy", Age: 88}) // 仅更新非零值字段

	// ** delete 删除语句
	db.Delete(&user, 1)

}
