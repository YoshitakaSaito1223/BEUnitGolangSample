package handler

import (
	"fmt"
	"net/http"
)

func HogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hoge")
}
