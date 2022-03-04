package main

import (
	"fmt"
)

var PostById map[int]*Post           // Idをキー，Post型の構造体をバリューとしたマップ
var PostsByAuthor map[string][]*Post // 著者名をキー，Post型の構造体を要素としたスライスをバリューとしたマップ

func store(post Post) {
	PostById[post.Id] = &post                                              // PostByIdに構造体Postのポインタを格納
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post) // PostsByAuthorのスライスに構造体Postのポインタを格納
}

func main1() {
	PostById = make(map[int]*Post)           // 空のマップ
	PostsByAuthor = make(map[string][]*Post) // 空のマップ

	// データ作成
	post1 := Post{Id: 1, Content: "Hello Piyo.", Author: "Be3"}
	post2 := Post{Id: 2, Content: "Hello Joh.", Author: "Be3"}
	post3 := Post{Id: 3, Content: "Piyoyo.", Author: "Be4"}
	post4 := Post{Id: 4, Content: "Piyoyoyoyo.", Author: "Be4"}

	// データ格納
	store(post1)
	store(post2)
	store(post3)
	store(post4)

	// データの確認
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Be3"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["Be4"] {
		fmt.Println(post)
	}
}
