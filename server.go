package main

import (
	"net/http"
	"simple_web/handlers" // go.modに指定したモジュール名 + パッケージ名

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()                 // マルチプレクサの生成
	mux.GET("/hello/:name", handlers.Hello) // 名前付きパラメータのあるURLとハンドラーHelloを対応
	mux.GET("/place_info/:place", handlers.PlaceInfo)
	// http.HandleFunc("/info", handlers.InfoHandler)
	mux.POST("/process", handlers.Process)               // 画像の送信先
	mux.GET("/status501", handlers.Status501)            // ステータスコード501とメッセージを返す
	mux.GET("/redirect_google", handlers.RedirectGoogle) // ステータスコード302でGoogleにリダイレクト
	mux.GET("/json", handlers.ResonseJson)               // JSONデータをレスポンス
	mux.GET("/set_cookie", handlers.SetCookie)           // クライアントにクッキーを保存
	mux.GET("/get_cookie", handlers.GetCookie)           // クライアントからクッキーを取得

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
