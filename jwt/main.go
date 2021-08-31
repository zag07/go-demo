package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/getToken", getToken)
	http.HandleFunc("/verify", verify)
	log.Fatal(http.ListenAndServe("localhost:8153", nil))
}

func getToken(w http.ResponseWriter, r *http.Request) {
	token, err := CreateToken("zs")
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
	}

	fmt.Fprintf(w, "Authorization = %q\n", token)
}

func verify(w http.ResponseWriter, r *http.Request) {
	_, err := ParseToken(r.Header.Get("Authorization"))
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
	}

	fmt.Fprintf(w, "success")
}
