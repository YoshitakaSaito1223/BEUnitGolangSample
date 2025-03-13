package handler

import (
	"fmt"
	"net/http"
)

func FugaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "FugaFuga")
}
