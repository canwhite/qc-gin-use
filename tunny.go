package main

import (
	"fmt"
	"time"

	"github.com/Jeffail/tunny"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CalculatorParams 计算器参数
type CalculatorParams struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
}

func TannyTest() {
	r := gin.Default()
	numCPUs := 4 // 设置池的大小为4
	p := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		params := payload.(CalculatorParams) // 将传入的任务数据转换为CalculatorParams类型
		sum := params.A + params.B           // 计算加法
		product := params.A * params.B       // 计算乘法
		time.Sleep(3 * time.Second)          // 模拟耗时操作
		return fmt.Sprintf("%d + %d = %d, %d * %d = %d", params.A, params.B, sum, params.A, params.B, product)
	})

	defer p.Close() // 在程序退出前关闭池
	r.POST("/calculator", func(c *gin.Context) {
		var params CalculatorParams
		// 绑定请求体的数据，并缓存到上下文中
		// 这里没有用ShouldBindJSON，是害怕请求请求体中的数据被一次消耗掉，所以用ShouldBindBodyWith
		/**

		BindJSON和ShouldBindJSON的区别是，BindJSON会在绑定失败时，
		自动返回一个400的状态码和一个错误信息给客户端，
		而ShouldBindJSON则不会，它只会返回一个错误值给调用者。

		如果你想自定义你的错误处理逻辑，你可以使用ShouldBindJSON方法，
		并根据返回的错误值来判断是否绑定成功。
		如果你想使用gin的默认错误处理逻辑，你可以使用BindJSON方法，
		它会帮你处理绑定失败的情况。
		*/

		if err := c.ShouldBindBodyWith(&params, binding.JSON); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		result, err := p.ProcessTimed(params, 10*time.Second) // 向池提交任务，并设置超时时间为10秒
		if err != nil {
			c.JSON(408, gin.H{
				"error": "timeout",
			})
			return
		}
		c.JSON(200, gin.H{ // 返回结果给客户端
			"result": result,
		})
	})
	r.Run(":8080")
}
