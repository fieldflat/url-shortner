package user_session

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redis_host     = "user_session_db"
	redis_port     = 6379
	redis_addr     = "user_session_db:6379"
	redis_password = ""
	redis_dbname   = 0
)

var ctx = context.Background()

type UserSessionRedisDB struct {
	UserSession *redis.Client
}

var redis_instance *UserSessionRedisDB

func GetUserSessionRedisDB() *UserSessionRedisDB {
	if redis_instance == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     redis_addr,
			Password: redis_password,
			DB:       redis_dbname, // use default DB
		})
		redis_instance = &UserSessionRedisDB{UserSession: client}
		log.Println("Successfully connected!")
	}
	log.Println("[GetUserSessionRedisDB] redis_instance: ", redis_instance)
	return redis_instance
}

func (u *UserSessionRedisDB) SetExpirationTime(sid string) error {
	err := u.UserSession.Set(ctx, sid, time.Now().Add(1*time.Hour), 0).Err()
	if err != nil {
		log.Println("[SetExpirationTime Set] err: ", err)
	}

	val, err := u.UserSession.Get(ctx, sid).Result()
	if err != nil {
		log.Println("[SetExpirationTime Get] err: ", err)
	}
	log.Println("sid: ", sid, ", value: ", val)
	return err
}

func (u *UserSessionRedisDB) IsExpired(sid string) bool {
	// _, exist := u.UserSession[sid]
	// return !exist || u.UserSession[sid].Before(time.Now())
	return false
}
