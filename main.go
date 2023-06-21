package main

import (
	"github.com/gin-gonic/gin"

	"goSeckill/config"
	"goSeckill/server"

)

func init(){

	config.InitConfig()

}

func main(){
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 路由组1 ，处理GET请求
	v1 := r.Group("/api")

	{
		v1.GET("/buyGood",server.BuyGood)
	//	v1.GET("submit", submit)
	}

	r.Run(":8000")
}
