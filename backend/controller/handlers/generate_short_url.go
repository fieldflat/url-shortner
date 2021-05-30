package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/fieldflat/url-shortner/backend/db/url_pairs"
	"github.com/fieldflat/url-shortner/backend/db/user_session"
	"github.com/fieldflat/url-shortner/backend/model/url"
	"github.com/google/uuid"
)

func GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	var sid string

	// cookie set
	if err != nil {
		u, err := uuid.NewRandom()
		if err != nil {
			log.Fatal("err when getting uuid: ", err)
			return
		}
		sid = u.String()
		cookie = &http.Cookie{
			Name:  "token",
			Value: sid,
		}
		http.SetCookie(w, cookie)

		db := user_session.GetUserSessionInMemoryDB()
		db.SetExpirationTime(sid)
	} else {
		sid = cookie.Value
	}

	// calcurate hash number
	originURL := exploitOriginURL(r.URL.Path)
	originURLInfo := url.OriginURLInfo{OriginURL: originURL, SID: sid}
	hash := hashFunc(originURL, cookie.Value)

	// store url_pair information
	db := url_pairs.GetURLPairsPostgresDB()
	db.SetURLPairs(url.ShortURL(hash), originURLInfo)
	prefix := "http://localhost:8080/re/" // TODO
	fmt.Fprintf(w, prefix+hash)
}

func exploitOriginURL(path string) string {
	str := "/short-url/" // TODO: get str from main.go
	return path[len(str):]
}

func hashFunc(originURL string, cookie string) string {
	hash := sha256.Sum256([]byte(originURL + cookie))
	EncodedHashStr := hex.EncodeToString(hash[:])
	len := calculateHashLength(originURL)
	return EncodedHashStr[:len]
}

func calculateHashLength(hashStr string) int {
	return len(hashStr) / 4
}
