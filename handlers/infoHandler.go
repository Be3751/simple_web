package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Post struct {
	User    string
	Threads []string
}

func InfoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Access to /info")

	fmt.Println("start to read a file")
	f, err := os.Open("templates/info.html")
	if err != nil {
		fmt.Println("error")
	}
	defer func() {
		fmt.Println("finish reading a file")
		f.Close()
	}()

	b, err := ioutil.ReadAll(f)
	s := string(b)
	io.WriteString(w, s)
}

func Hello(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func PlaceInfo(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "you are at %s!\n", p.ByName("place"))
}

// リクエストから受け取ったファイルをバイト配列にしてレスポンス
func Process(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	file, _, err := req.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func Status501(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(501) // ステータスコードを501に設定
	fmt.Fprintf(w, "No such a service.")
}

func RedirectGoogle(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// WriteHeaderの呼び出し後にヘッダの変更はできないため，WriteHeaderの呼び出し順序に注意．
	w.Header().Set("Location", "http://google.com") // レスポンスヘッダにLocation: http://google.comを設定
	w.WriteHeader(302)                              // ステータスコードを302に設定
}

func ResonseJson(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json") // ヘッダにjsonをレスポンスすることを設定
	// 構造体Postを複合リテラル定義
	post := &Post{
		User:    "Be3",
		Threads: []string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post) // json.Marshalで構造体をJSON形式にエンコード
	w.Write(json)                 // メッセージに書き込み
}
