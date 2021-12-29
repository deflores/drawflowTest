package main

import (
	"app/drawflow"
	"net/http"
)

func main() {
	r := drawflow.Setup()

	http.ListenAndServe(":80", r)
}
