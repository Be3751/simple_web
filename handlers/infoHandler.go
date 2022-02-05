package handlers

import (
	"io"

	"net/http"
)

func InfoHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>Here is the /info endpoint.</h1>
	</body>
	</html>
	`)

}
