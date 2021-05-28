package handlers

import (
	"net/http"
	"strings"
)

func RedirectOriginURL(w http.ResponseWriter, r *http.Request) {
	ok := isRootAccess(r.URL.Path)
	if !ok {
		http.NotFound(w, r)
		return
	}
	shortURL := r.URL.Path[1:]
	originURL := getOriginURL(shortURL)
	http.Redirect(w, r, originURL, http.StatusSeeOther)
}

func getOriginURL(shortURL string) string {
	return "https://google.com"
}

func isRootAccess(path string) bool {
	return !strings.HasPrefix(path, "short-url")
}
