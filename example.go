package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

//- 모델
//- id
//- title
//- category(nullable)
//- description(nullable)
//- created_at
//- updated_at

type Basic struct {
	ID        uint           `gorm:"primaryKey;autoIncrementIncrement" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Post struct {
	//gorm.Model // === ID, CreatedAt, UpdatedAt을 기본적으로 가지고 있음 === 본인이 만든 BasicPost랑 같은 기능
	//Basic       Basic   `gorm:"embedded"`
	ID          uint           `gorm:"primaryKey;autoIncrementIncrement" json:"id"`
	CreatedAt   time.Time      `gorm:"autoCreateTime:nano" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Title       string         `gorm:"not null" json:"title"'`
	Category    *string        `json:"category"` //sql.NullString은 구조체 타입으로 JSON값을 해당 타입으로 변경할 수 없음. nullable하게 하려면 *string 사용
	Description *string        `json:"description"`
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
		panic("db ERROR: (initDatabase)")
	}
	err = db.AutoMigrate(
		&Post{},
	)
	if err != nil {
		fmt.Println("Failed AutoMigrate")
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
	reqBody := Post{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post Post
	post.Title = reqBody.Title
	post.Category = reqBody.Category
	post.Description = reqBody.Description
	// HTTP 요청 바디에서 JSON 읽어오기
	if result := DB.Create(&post); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusCreated, post)
}

func GetAllPost(c *gin.Context) {
	rows, err := DB.Raw("SELECT * FROM posts").Rows()
	if err != nil {
		//	Error Handling
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post
		// 구조체의 field순서까지 맞춰줘야한다니..
		err := rows.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.Title, &post.Category, &post.Description, &post.DeletedAt)
		if err != nil {
			// 에러 처리
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		posts = append(posts, post)
	}
	fmt.Println(posts)
	if err := rows.Err(); err != nil {
		// 에러 처리
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"data": posts,
		})
}

func DeletePost(c *gin.Context) {
	var post Post
	postId := c.Params.ByName("id")
	if err := DB.Raw("DELETE FROM posts WHERE id = ?", postId).Scan(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"id": postIdInt})
}

func setUpRouter() *gin.Engine {
	r := gin.Default() // gin에서 기본 라우터를 담당
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3001"},
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
	}))

	r.GET("/post/:id", GetPost)
	r.DELETE("/post/:id", DeletePost)
	r.GET("/posts", GetAllPost)
	r.POST("/post", CreatePost)
	return r
}

func main() {

	initDatabase()
	r := setUpRouter()

	// ! Question : 왜 안되는지 모르겠다.
	//err := DB.Raw("INSERT INTO posts (title) VALUES (?)", "Hello")
	//if err != nil {
	//	panic("INSERT ERROR")
	//}

	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}
