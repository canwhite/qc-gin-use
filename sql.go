package main

import (
	"fmt"
)

func TestMysql() {

	//然后将mysql植进来,配置dsn属性
	//名称:密码@tcp(url:port)/
	//dsn := "root:715705@Qc123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True"
	// //然后用dsn和渠道，拿到db
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db := GetInstance()

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
	db.First(&user, "name = ?", "zhagnsan") // 查找 name 字段值为 linzy 的记录
	//db.Where("name = ?", "jinzhu").First(&user) // 根据条件查找第一条记录1
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
	db.Model(&user).Updates(User{Name: "zhagnsan", Age: 88}) // 仅更新非零值字段

	// ** delete 删除语句
	// db.Delete(&user, 1)
	// 使用Where方法指定查询条件
	//db.Where("name = ?", "Sub Menu A3").Delete(&SubMenu{})
	//db.Delete(&SubMenu{}, "name = ?", "Sub Menu A3")

}

func ForeignKeyTest() {
	db := GetInstance()

	// 自动迁移表结构
	db.AutoMigrate(&Menu{}, &SubMenu{})

	//创建一些数据
	db.Create(&Menu{
		Name: "Menu A",
		SubMenus: []SubMenu{
			{Name: "Sub Menu A1"},
			{Name: "Sub Menu A2"},
		},
	})

	db.Create(&Menu{
		Name: "Menu B",
		SubMenus: []SubMenu{
			{Name: "Sub Menu B1"},
			{Name: "Sub Menu B2"},
		},
	})

	//查询所有的一级菜单和二级菜单

	//查询名字为"Menu A"的一级菜单和二级菜单
	var menuA Menu
	db.Preload("SubMenus").Where("name = ?", "Menu A").First(&menuA)
	fmt.Println("---a---", menuA)

	var menus []Menu
	//find发现所有，First匹配一个
	db.Preload("SubMenus").Find(&menus)
	fmt.Println(menus)

	//给MenuA添加二级菜单
	// 给"Menu A"添加一个二级菜单"Sub Menu A3"
	db.Model(&menuA).Association("SubMenus").Append(&SubMenu{Name: "Sub Menu A3"})

	// 删除"Menu A"的二级菜单"Sub Menu A1"
	db.Model(&menuA).Association("SubMenus").Delete(&SubMenu{Name: "Sub Menu A1"})

	//当然可以update
	//db.Model(&SubMenu{}).Where("name = ?", "Sub Menu A3").Update("name", "Sub Menu A4")
	db.Model(&menuA).Association("SubMenus").Replace(&SubMenu{Name: "Sub Menu A3"}, &SubMenu{Name: "Sub Menu A6"})

}

// 写一个连接查询
func JoinQueryTest() {

	//连接查询的时候，连接一级菜单和二级菜单的名字
	db := GetInstance()
	var list []struct {
		MenuName    string
		SubMenuName string
	}
	/**
	- Table:选择数据表
	- Joins:添加表连接
	- Select:选择需要的列,可以使用别名
	- Scan:执行查询并扫描结果至变量
	*/
	db.Table("menus").Joins("JOIN sub_menus ON sub_menus.menu_id = menus.id").Select("menus.name as menu_name, sub_menus.name as sub_menu_name").Scan(&list)
	fmt.Println("list", list)
}

//写一个聚合查询

// 测试下子查询
func SubQueryAndGroupTest() {
	// SELECT * FROM sub_menus WHERE name IN (SELECT name FROM menus WHERE name LIKE 'Menu%');

	db := GetInstance()
	// 子查询一级菜单的名字
	var results []User
	//Table可以放在前边也可以放在后边
	subQuery := db.Table("users").Select("AVG(age)").Where("name LIKE ?", "name%")
	db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
	fmt.Println(results)

}
