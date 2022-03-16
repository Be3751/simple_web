// GORMを用いたDB操作
// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/lib/pq"
// )

// // 投稿とコメントで1対多の関係になるようフィールドにポインタを追加
// // 投稿のための構造体
// type Post struct {
// 	Id        int
// 	Content   string
// 	Author    string `sql:"not null"`
// 	Comments  []Comment
// 	CreatedAt time.Time
// }

// // コメントのための構造体
// type Comment struct {
// 	Id        int
// 	Content   string
// 	Author    string `sql:"not null"`
// 	PostId    int
// 	CreatedAt time.Time
// }

// var Db *gorm.DB // 構造体sql.DBの宣言

// // 構造体sql.DBの初期化
// func init() {
// 	var err error
// 	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=ppp sslmode=disable") // DB接続
// 	if err != nil {
// 		panic(err)
// 	}
// 	Db.AutoMigrate(&Post{}, &Comment{})
// }

// func main() {
// 	post := Post{Content: "Hello, Piyo!", Author: "Be3"}
// 	fmt.Println(post)

// 	Db.Create(&post) // Gormによりレコードの作成
// 	fmt.Println(post)

// 	comment := Comment{Content: "いい投稿だね！", Author: "Be3"}
// 	Db.Model(&post).Association("Comments").Append(comment)

// 	var readPost Post
// 	Db.Where("author = $1", "Be3").First(&readPost)
// 	var comments []Comment
// 	Db.Model(&readPost).Related(&comments)
// 	fmt.Println(comments[0])
// }