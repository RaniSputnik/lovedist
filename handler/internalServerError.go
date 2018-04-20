package handler

import (
	"fmt"
	"net/http"
)

// TODO move to common template
const internalServerErrorHTML = `<!doctype html>
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

		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.0/css/bootstrap.min.css" integrity="sha384-9gVQ4dYFwwWSjIDZnLEWnxCjeSWFphJiwGPXr1jddIhOegiu1FwO5qRGvFXOdJZ4" crossorigin="anonymous">

		<style>
			.container {
				max-width: 520px;
			}

			body {
				padding-bottom: 100px;
			}

			.footer {
				position: fixed;
				bottom: 0;
				width: 100%;
				height: 60px;
				line-height: 60px;
				background-color: #f5f5f5;
			}
		</style>
	</head>
	<body>
		<div class="container text-center">
			</br><h1>Something went wrong</h1></br>
			
			<p class="lead">Unfortunately there was a problem with the server, please try again later.</p>
			<a class="btn btn-primary" href="/">Return to home</a>
		</div>

		<footer class="footer">
			<div class="container text-center">
				<span class="text-muted"><small>Made with LÖVE by <a href="http://ryanloader.me">Ryan Loader<a/></small></span>
			</div>
		</footer>
	</body>
</html>
`

func internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println("Internal Server Error:", err)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(internalServerErrorHTML))
}
