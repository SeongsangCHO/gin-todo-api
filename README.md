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