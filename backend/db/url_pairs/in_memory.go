package url_pairs

import (
	"log"

	"github.com/fieldflat/url-shortner/backend/db/user_session"
	"github.com/fieldflat/url-shortner/backend/model/url"
)

type URLPairsDB struct {
	URLPairs map[url.ShortURL]url.OriginURLInfo
}

var instance *URLPairsDB

func GetURLPairsDB() *URLPairsDB {
	if instance == nil {
		instance = &URLPairsDB{URLPairs: make(map[url.ShortURL]url.OriginURLInfo)}
	}
	log.Println("[GetURLPairsDB] instance = ", instance)
	return instance
}

func (u *URLPairsDB) SetURLPairs(shortURL url.ShortURL, originURLInfo url.OriginURLInfo) error {
	u.URLPairs[shortURL] = originURLInfo
	// TODO: Error Check if neccesary
	log.Println("[SetURLPairs] instance = ", instance)
	return nil
}

func (u *URLPairsDB) GetOriginURLInfo(shortURL url.ShortURL) url.OriginURLInfo {
	return u.URLPairs[shortURL]
}

func (u *URLPairsDB) IsExpired(shortURL url.ShortURL) bool {
	sid := u.URLPairs[shortURL].SID
	instance := user_session.GetUserSessionInMemoryDB()
	b := instance.IsExpired(sid)
	return b
}
