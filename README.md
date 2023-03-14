# 마샬링 (Marshaling)
- 메모리에 올라간 값을 정수 또는 다른 타입으로 변환하는 과정
- go에선 정수형, 구조체같은 타입을 byte slice로 변경하는 것
- json.Marshal함수가 해당 역할을 해줌
```go
type Post struct{
	Title string
}

func main() {
	var post = Post{Title:"82"}
	bytes, _ := json.Marshal(post)
	fmt.Println(bytes) // [11, 22, 33, 44....
	fmt.Println(string(bytes)) // {"Title": "82"} 
	//Title을 소문자로 하려면 Title string `json:"title"`로 구조체 타입정의
	//이외 인덴트를 추가하거나, 특정 값을 제외하거나 하는 등의 마샬링 제어가 가능하다.
}
```

# 언마샬링
- byte slice나 문자열을 논리적 자료구조로 변경하는 것을 언마샬링
```go
var t bool
json.Unmarshal([]byte("true"), &t) // byte slice, 주소로 인자 전달
fmt.Printf("%t\n", b) //true
```


# GORM 연결
- postgres install
```
https://www.postgresql.org/download/ 에서 버전설치 //현재 프로젝트 15버전
```

- `pgAdmin` or psql로 database 생성
-  gorm 연결코드로 생성되어있는 DB와 gorm연결
```go
db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1234 dbname=gorm-todo-api port=5432 sslmode=disable TimeZone=Asia/Seoul",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
```