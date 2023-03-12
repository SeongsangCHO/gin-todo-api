package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//gin.Context -> HTTP요청과 관련된 정보를 갖고있는 구조체, 요청에 대한 처리를 담당하는 핸들러함수
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
