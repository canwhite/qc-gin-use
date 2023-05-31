package main

import (
	"context"

	"github.com/gin-gonic/gin"
	//加上v8相当于在使用最新版本，可以用最新的api
	"github.com/go-redis/redis/v8"
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
	r.GET("/ping", pingHandler)
	//牛逼点的get , 假设有一个get路由 /user/:name?age=18
	r.GET("/user/:name", getHandler)

	//--post
	r.POST("/post", postHandler)

	//--redis的get和set
	r.GET("/set/:key/:value", rsetHandler)

	r.GET("/get/:key", rgetHandler)

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
func pingHandler(c *gin.Context) {
	c.String(200, "pong")
}

// 牛逼点的get
// 假设有一个get路由 /user/:name?age=18
func getHandler(c *gin.Context) {
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
func postHandler(c *gin.Context) {
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

func rsetHandler(c *gin.Context) {

	//开始吧
	key := c.Param("key")
	value := c.Param("value")

	//然后我们来存它们

	/** 参数解析
	* context.Background(): 这是一个context.Context类型的参数，用于传递一些上下文信息，比如超时、取消、截止时间等。context.Background()是一个空的context，没有任何信息。
		它是所有其他context的根节点
		所以redis实际上依赖了go本身的context，用于处理请求的超时、取消等行为
	* key: 这是一个string类型的参数，用于指定要设置的键的名称。
	* value: 这是一个interface{}类型的参数，用于指定要设置的键的值。它可以是任何类型，go-redis会自动将其转换为字符串。
	* 0: 这是一个time.Duration类型的参数，用于指定要设置的键的过期时间。如果为0，表示永不过期。
	* Err(): 这是一个方法，用于返回rdb.Set操作的错误信息。如果没有错误，返回nil。
	*/
	err := rdb.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "ok")
}

// 然后是redis的set和get handle操作
func rgetHandler(c *gin.Context) {
	key := c.Param("key")
	value, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, value)
}
