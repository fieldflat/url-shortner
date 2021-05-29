package user_session

import (
	"log"
)

type UserSessionInterface interface {
	SetExpirationTime(sid string) error
	IsExpired(sid string) bool
}

func SetUserSession(u UserSessionInterface, sid string) {
	err := u.SetExpirationTime(sid)
	if err != nil {
		log.Println("err: ", err)
	}
}

func ValidateUserSession(u UserSessionInterface, sid string) bool {
	return u.IsExpired(sid)
}
