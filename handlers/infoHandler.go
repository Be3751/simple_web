package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

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
