package main

import (
	"html/template"
	"net/http"
)

const indexHTML = `<!doctype html>
<html class="no-js" lang="">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="x-ua-compatible" content="ie=edge">
		<title>LÖVE Dist</title>
		<meta name="description" content="">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

		<link rel="manifest" href="site.webmanifest">
		<link rel="apple-touch-icon" href="icon.png">
		<!-- Place favicon.ico in the root directory -->

		<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/normalize/8.0.0/normalize.min.css">
	</head>
	<body>
		<form enctype="multipart/form-data" action="http://127.0.0.1:9090/upload" method="post">
			<input type="file" name="uploadfile" />
			<input type="hidden" name="token" value="{{.}}"/>
			<input type="submit" value="upload" />
		</form>
	</body>
</html>
`

func indexHandler() http.HandlerFunc {
	t := template.Must(template.New("index").Parse(indexHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
