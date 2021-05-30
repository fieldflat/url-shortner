package user_session

// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type UserSessionPostgresDB struct {
	UserSession *sql.DB
}

var psql_instance *UserSessionPostgresDB

func GetUserSessionPostgresDB() *UserSessionPostgresDB {
	if psql_instance == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		log.Println("Successfully connected!")
		psql_instance = &UserSessionPostgresDB{UserSession: db}
	}
	log.Println("[GetUserSessionPostgresDB] psql_instance: ", psql_instance)
	return psql_instance
}

func (u *UserSessionPostgresDB) SetExpirationTime(sid string) error {
	var id string
	sqlStatement := `
INSERT INTO user_session (sid, expiration)
VALUES ($1, $2)`
	err := u.UserSession.QueryRow(sqlStatement, sid, time.Now().Add(1*time.Hour)).Scan(&id)
	if err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

func (u *UserSessionPostgresDB) IsExpired(sid string) bool {
	// _, exist := u.UserSession[sid]
	// return !exist || u.UserSession[sid].Before(time.Now())
	return false
}
