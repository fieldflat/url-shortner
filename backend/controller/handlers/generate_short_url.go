package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

func GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	originURL := exploitOriginURL(r.URL.Path)
	hash := hashSha256(originURL)
	fmt.Fprintf(w, hash)
}

func exploitOriginURL(path string) string {
	str := "/short-url/" // TODO: get str from main.go
	return path[len(str):]
}

func hashSha256(originURL string) string {
	hash := sha256.Sum256([]byte(originURL))
	EncodedHashStr := hex.EncodeToString(hash[:])
	len := calculateHashLength(EncodedHashStr)
	return EncodedHashStr[:len]
}

func calculateHashLength(hashStr string) int {
	return len(hashStr) / 4
}
