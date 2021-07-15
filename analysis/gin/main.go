package main

import (
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
