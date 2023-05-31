package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// 定义redis客户端
var rdb *redis.Client

func TestRedis() {
	rdb = redis.NewClient(&redis.Options{
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//创建gin引擎
	r := gin.Default()
	//--简单get
	r.GET("/ping", pingHandle)
	//牛逼点的get , 假设有一个get路由 /user/:name?age=18
	r.GET("/user/:name", getHandle)

	//--post
	r.POST("/post", postHandle)

	//--redis的get和set

	// 启动服务器
	r.Run(":8080")
}

/********************************************************

pre:Context

*接值

--post和get ：c.Param("name")和c.Query("age")。
--c.Header("Content-Type", "application/json")和c.JSON(200, gin.H{"message": "ok"})
--本身也可以管理state，c.Set("key", "value")和c.Get("key")。

* 返回值
--返回json // 在gin.Context中返回json
c.JSON(200, user) // 返回 {"name": "Alice", "age": 18}
********************************************************/

// 简单点的get
func pingHandle(c *gin.Context) {
	c.String(200, "pong")
}

// 牛逼点的get
// 假设有一个get路由 /user/:name?age=18
func getHandle(c *gin.Context) {
	// 从路径中获取name参数
	name := c.Param("name")
	// 从查询字符串中获取age参数
	age := c.Query("age")

	// 返回json响应
	//gin.H 是一个 map[string]interface{} 的简写,用于快速构造 JSON 对象。
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// post操作
func postHandle(c *gin.Context) {
	//从json中绑定数据到Params的结构体
	var params Params
	err := c.BindJSON(&params)
	//如果有报错，收一下边
	if err != nil {
		//处理一下错误
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	//如果没有错误，我们接着往下走就完事了
	c.JSON(200, gin.H{
		"name":  params.name,
		"email": params.email,
	})

}

//然后是redis的set和get handle操作
