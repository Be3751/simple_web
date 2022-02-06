package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
