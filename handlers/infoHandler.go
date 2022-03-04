package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

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

// セッションクッキーはブラウザを終了すると破棄される．
// タブやウィンドウを閉じるだけでは破棄されない．
func SetCookie(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// Cookie構造体を定義
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Piyo is sleeping",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Masa is eating",
		HttpOnly: true,
	}
	// ヘッダにシリアル化したクッキーを設定
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func GetCookie(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// 本来であればクッキーの解析を自前で行う必要があるが，Go言語に用意されている
	c1, err := req.Cookie("first_cookie") // キーの値で一意に取得する方法
	if err != nil {
		fmt.Fprintln(w, "Cannot get the cookie.")
	}
	cs := req.Cookies() // 全ての組を取得する方法
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

// フラッシュメッセージにしようするメッセージをクッキーに設定
func SetMessage(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg), // メッセージに特殊文字が含まれる可能性を考慮して，通常URLエンコードする必要がある
	}
	http.SetCookie(w, &c)
}

// フラッシュメッセージ
func ShowMessage(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	c, err := req.Cookie("flash") // クッキーからflashがキーとなる値を取得
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message.")
		}
	} else {
		// 同じ名前のクッキーに過去を表す値を指定することで，実質的に破棄している
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value) // クッキーflashを破棄する前に取得したメッセージをデコード
		fmt.Fprintln(w, string(val))
		fmt.Fprintln(w, "hello")
	}
}

func ShowInfo(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("templates/info.html")
	t.Execute(w, "Hello")
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

// テンプレートエンジンの起動
func ProcessTemplate(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("templates/practice.html").Funcs(funcMap)
	t, _ = t.ParseFiles("templates/practice.html")
	t.Execute(w, time.Now())
}

func Layout(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("templates/layout.html", "templates/red_hello.html")
	} else {
		t, _ = template.ParseFiles("templates/layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
