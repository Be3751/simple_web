package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// 投稿とコメントで1対多の関係になるようフィールドにポインタを追加
// 投稿のための構造体
type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

// コメントのための構造体
type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

var Db *sql.DB // 構造体sql.DBの宣言

// 構造体sql.DBの初期化
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=ppp sslmode=disable") // DB接続
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	rows, err := Db.Query("select id, content, author from comments where post_id = $1", id) // 外部キーpost_idに一致するコメントを複数取得（1対多の関係）
	if err != nil {
		return
	}
	// 取得したコメントを投稿内のコメントスライスに追加していく
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func (comment *Comment) Create() (err error) {
	// コメントをDBに保存する前に投稿が存在していることを確認する
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func main() {
	post := Post{Content: "Hello, Piyo!", Author: "Be3"}

	fmt.Println(post)
	post.Create()     // レコードの作成
	fmt.Println(post) //  レコードを作成し，Idに値があることを確認

	// 単一のレコード取得が可能であることを確認
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	// レコードの更新
	readPost.Content = "Hello, Be3!"
	readPost.Author = "Piyo"
	readPost.Update()

	// 上限10個までのレコードをスライスで所得
	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete() // レコードの削除
}
