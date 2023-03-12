package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

//- 모델
//- id
//- title
//- category(nullable)
//- description(nullable)
//- created_at
//- updated_at

type BasicPost struct {
	ID        uint      `gorm:"primaryKey;autoIncrementIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
}

type Post struct {
	//gorm.Model === ID, CreatedAt, UpdatedAt을 기본적으로 가지고 있음 === 본인이 만든 BasicPost랑 같은 기능
	BasicPost   BasicPost `gorm:"embedded"`
	Title       string    `gorm:"not null"`
	Category    sql.NullString
	Description sql.NullString
}

func main() {
	// gorm은 기본적으로 statement cache를 사용하고 있음(SQL문을 캐싱하여 재사용)
	// PreferSimpleProtocol은 statement cache를 대신, 간단한 쿼리요청을 사용하는 것 -> 두 옵션을 비교하여 적절히 사용하는 것을 권장함
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1234 dbname=gorm-todo-api port=5432 sslmode=disable TimeZone=Asia/Seoul",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("DB 연결에 실패했습니다")
	}
	post := Post{Title: "HEllo"}
	db.Create(&post)

	r := gin.Default()
	//gin.Context -> HTTP요청과 관련된 정보를 갖고있는 구조체, 요청에 대한 처리를 담당하는 핸들러함수
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/post", func(c *gin.Context) {
		//	모든 post를 return하기
	})

	r.POST("/post", func(c *gin.Context) {
		//	POST 생성하기
	})

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
