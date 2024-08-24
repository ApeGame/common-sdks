package main

import (
	_ "github.com/ApeGame/common-sdks/api"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		panic(err)
	}
}
