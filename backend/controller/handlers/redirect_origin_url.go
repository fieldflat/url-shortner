package handlers

import (
	"net/http"

	"github.com/fieldflat/url-shortner/backend/db/url_pairs"
	"github.com/fieldflat/url-shortner/backend/model/url"
)

func RedirectOriginURL(w http.ResponseWriter, r *http.Request) {
	str := "/re/" // TODO: get str from main
	shortURL := r.URL.Path[len(str):]
	originURL := getOriginURL(shortURL)
	http.Redirect(w, r, originURL, http.StatusSeeOther)
}

func getOriginURL(shortURL string) string {
	db := url_pairs.GetURLPairsDB()
	return db.GetOriginURLInfo(url.ShortURL(shortURL)).OriginURL
}
