package main

import "github.com/gin-gonic/gin"

/**
 * @Author: guohuliu
 * @Email: liuguotiger@aliyun.com
 * @Date: 2020/11/14 10:35 上午
 * @Desc:
 */

func main() {
	r := gin.Default()
	r.GET("/echo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Run()
}
