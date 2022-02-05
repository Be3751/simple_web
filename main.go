package main

import (
	"net/http"

	"simple_web/handlers" // go.modに指定したモジュール名 + パッケージ名
)

func main() {
	http.HandleFunc("/info", handlers.InfoHandler)
	http.ListenAndServe("localhost:8000", nil)
}
