package main

import (
	"fmt"
	"net/http"

	"github.com/YoshitakaSaito1223/demo/handler"
)

func main() {
	fmt.Println("Welcome to Go World!")

	fmt.Println("listen and serve :60000")
	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/hoge", handler.HogeHandler)
	http.HandleFunc("/fuga", handler.FugaHandler)
	http.ListenAndServe(":60000", nil)
}
