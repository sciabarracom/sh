## Quickstart

	# start an http server
	python3 -m http.server

	# build
	GOOS=js GOARCH=wasm go build -o parser.wasm

	# point your browser at the http server
	firefox http://localhost:8000/

If the Go changes, rebuild.
