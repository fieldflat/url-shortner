package handlers

import (
	"fmt"
	"net/http"
)

func ReturnTopPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "hello")
		return
	}
}
