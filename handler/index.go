package handler

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
			</br><h1>LÖVE Distributor</h1></br>
			
			<p class="lead">Build and distribute your <a href="https://love2d.org/" target="_blank">LÖVE</a> games easily.</p>
			</br>

			<div class="card">
				<div class="card-body">
					<form class="form-inline" enctype="multipart/form-data" action="/build" method="POST">
						<div class="form-group">
							<input type="file" class="form-control-file" id="fileInput" name="uploadfile" accept=".love" required>
						</div>
						<button class="btn btn-primary" type="submit">Upload</button>
					</form>
				</div>
			</div>

			</br>

			<div class="card">
				<h5 class="card-header">How it works</h5>
				<ul class="list-group list-group-flush">
					<li class="list-group-item">Upload your .love file using the form above.</li>
					<li class="list-group-item">We'll send you executables for Windows, Mac and Linux.</li>
					<li class="list-group-item">Download and play the game or share with friends.</li>
				</ul>
			</div>
		</div>

		<footer class="footer">
			<div class="container text-center">
				<span class="text-muted"><small>Made with LÖVE by <a href="http://ryanloader.me">Ryan Loader<a/></small></span>
			</div>
		</footer>
	</body>
</html>
`

func indexHandler() http.HandlerFunc {
	t := template.Must(template.New("index").Parse(indexHTML))
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	}
}
