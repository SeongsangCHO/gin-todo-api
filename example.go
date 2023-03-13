package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
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

var DB *gorm.DB

func initDatabase() *gorm.DB {
	// gorm은 기본적으로 statement cache를 사용하고 있음(SQL문을 캐싱하여 재사용)
	// PreferSimpleProtocol은 statement cache를 대신, 간단한 쿼리요청을 사용하는 것 -> 두 옵션을 비교하여 적절히 사용하는 것을 권장함
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1234 dbname=gorm-todo-api port=5432 sslmode=disable TimeZone=Asia/Seoul",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("db ERROR: (initDatabase)", err)
	}
	DB = db
	return db
}

func GetPost(c *gin.Context) {
	var post Post
	postId := c.Params.ByName("id")
	if err := DB.Raw("SELECT * FROM posts WHERE id = ?", postId).Scan(&post).Error; err != nil {
		// TODO - > 값이 조회는 되는데 에러처리가 제대로 되지 않는 이슈가 있음 Scan이후 값이 없을때를 어떻게 catch하는지 ?
		// return 받은 db data를 직렬화해야하는데 그 방법은 어떻게 하는지?
		// db connection은 언제 어느 시점에 끊어야하는지 ?
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found!"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {

}

func setUpRouter() *gin.Engine {
	r := gin.Default() // gin에서 기본 라우터를 담당

	r.GET("/post/:id", GetPost)
	r.POST("/post", CreatePost)
	return r
}

func main() {

	initDatabase()
	r := setUpRouter()

	var post Post
	DB.Raw("SELECT id, title FROM posts WHERE id = ?", 1).Scan(&post)
	fmt.Printf("%+v\n", post)

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
