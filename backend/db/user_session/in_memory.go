package user_session

import (
	"log"
	"time"
)

type UserSessionInMemoryDB struct {
	UserSession map[string]time.Time
}

// singleton instance
var instance *UserSessionInMemoryDB

func GetUserSessionInMemoryDB() *UserSessionInMemoryDB {
	if instance == nil {
		instance = &UserSessionInMemoryDB{UserSession: make(map[string]time.Time)}
	}
	log.Println("[GetUserSessionInMemoryDB] instance: ", instance)
	return instance
}

func (u *UserSessionInMemoryDB) SetExpirationTime(sid string) error {
	u.UserSession[sid] = time.Now().Add(1 * time.Hour)
	log.Println("[SetExpirationTime] instance: ", instance)
	// TODO: Error check
	return nil
}

func (u *UserSessionInMemoryDB) IsExpired(sid string) bool {
	_, exist := u.UserSession[sid]
	return !exist || u.UserSession[sid].Before(time.Now())
}
