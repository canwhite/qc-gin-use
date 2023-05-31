package main // 定义包名

import (
	"database/sql" // 导入sql包，用于操作数据库
	"sync"         // 导入sync包，用于同步操作
	"time"         // 导入time包，用于设置时间

	"gorm.io/driver/mysql" // 导入gorm的mysql驱动包
	"gorm.io/gorm"         // 导入gorm包，用于封装数据库操作
)

/*
这样不算是起别名，而是定义了一个新的类型DB，它嵌入了gorm.DB类型作为一个匿名字段。
这样的好处是，DB类型可以直接调用gorm.DB类型的所有方法和属性，而不需要通过字段名来访问。
同时，DB类型也可以定义自己的方法和属性，来扩展或覆盖*gorm.DB类型的功能。
这是一种组合的方式，而不是继承的方式。
*/
type DB struct {
	*gorm.DB // 嵌入字段，表示DB类型继承了*gorm.DB类型的所有方法和属性
}

func (db *DB) Close() error {
	// 定义一个Close方法，用于关闭数据库连接
	sqlDB, err := db.DB.DB() // 调用嵌入字段的DB方法，获取底层的*sql.DB对象
	if err != nil {
		return err // 如果出错，返回错误
	}
	return sqlDB.Close() // 调用*sql.DB对象的Close方法，关闭数据库连接，并返回结果
}

var instance *DB   // 定义一个全局变量instance，用于存储单例对象
var once sync.Once // 定义一个全局变量once，用于保证只执行一次初始化操作

func GetInstance() *DB {
	// 定义一个GetInstance函数，用于获取单例对象，如果没有初始化则先初始化
	// once只执行一次，当然我们也可以写在最外层init里
	once.Do(func() {
		// 使用once.Do方法，传入一个匿名函数，保证只执行一次
		sqlDB, e := sql.Open("mysql", "root:715705@Qc123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
		// 调用sql.Open函数，传入驱动名和数据源名，打开一个数据库连接，并返回一个*sql.DB对象和一个错误值
		if e != nil {
			panic("sql open failed") // 如果出错，抛出panic异常
		}
		// 设置sql池的最大空闲连接数
		sqlDB.SetMaxIdleConns(10)
		// 设置sql池的最大打开连接数
		sqlDB.SetMaxOpenConns(100)
		// 设置sql池的连接过期时间
		sqlDB.SetConnMaxLifetime(time.Hour)
		// 调用gorm.Open函数，传入一个mysql驱动对象和一个gorm配置对象，打开一个gorm数据库连接，并返回一个*gorm.DB对象和一个错误值
		// mysql驱动对象是通过mysql.New函数创建的，传入一个mysql配置对象，其中指定了Conn为之前创建的*sql.DB对象
		// gorm配置对象是直接使用默认值的
		db, err := gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})

		if err != nil {
			panic("failed to connect database") // 如果出错，抛出panic异常
		}

		instance = &DB{db} // 将*gorm.DB对象赋值给instance变量，并将其转换为*DB类型
	})
	return instance // 返回instance变量作为结果
}
