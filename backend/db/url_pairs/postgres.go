package url_pairs

// https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fieldflat/url-shortner/backend/model/url"
	_ "github.com/lib/pq"
)

const (
	host     = "url_pairs_db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type URLPairsPostgresDB struct {
	URLPairs *sql.DB
}

var psql_instance *URLPairsPostgresDB

func GetURLPairsPostgresDB() *URLPairsPostgresDB {
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
		psql_instance = &URLPairsPostgresDB{URLPairs: db}
	}
	log.Println("[GetURLPairsPostgresDB] psql_instance = ", psql_instance)
	return psql_instance
}

func (u *URLPairsPostgresDB) SetURLPairs(shortURL url.ShortURL, originURLInfo url.OriginURLInfo) error {
	var id string
	sqlStatement := `
INSERT INTO url_pairs (short_url, origin_url, sid)
VALUES ($1, $2, $3)`
	err := u.URLPairs.QueryRow(sqlStatement, shortURL, originURLInfo.OriginURL, originURLInfo.SID).Scan(&id)
	if err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

func (u *URLPairsPostgresDB) GetOriginURLInfo(shortURL url.ShortURL) url.OriginURLInfo {
	sqlStatement := `
SELECT origin_url, sid FROM url_pairs WHERE short_url = $1 LIMIT 1`
	row, err := u.URLPairs.Query(sqlStatement, shortURL)
	if err != nil {
		log.Println("err: ", err)
		return url.OriginURLInfo{}
	}

	var origin_url string
	var sid string
	row.Next()
	row.Scan(&origin_url, &sid)
	return url.OriginURLInfo{OriginURL: origin_url, SID: sid}
}

func (u *URLPairsPostgresDB) IsExpired(shortURL url.ShortURL) bool {
	return false
}

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected!")
// }
