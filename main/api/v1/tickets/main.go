package tickets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// структура соединения с базой данных
type dbConn struct {
	user string
	pass string
	base string
	conn *sql.DB
}

// метод коннекта к базе
func (db *dbConn) connect() {
	conn, _ := sql.Open("mysql", db.user+":"+db.pass+"@/"+db.base)
	db.conn = conn
}

func getConnection() *sql.DB {
	db := dbConn{
		user: "goexp",
		pass: "F76MWx3Px2",
		base: "exp_golang",
	}
	db.connect()
	return db.conn
}

func GetList(w http.ResponseWriter, r *http.Request) {

	file, _ := ioutil.ReadFile("/storage/data.json")
	storage := data{}

	_ = json.Unmarshal([]byte(file), &storage)

	_, err := fmt.Fprintf(w, "Hello, my name is %s %s!", storage.Name, storage.LastName)
	if err != nil {
		return
	}

}

func Read(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {

	db := getConnection()

	_, createErr := db.Query("CREATE TABLE `tickets` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `url` varchar(50) NOT NULL)")
	if createErr != nil {
		return
	}

}

func Add(w http.ResponseWriter, r *http.Request) {

	db := getConnection()
	_, createErr := db.Query("INSERT INTO tickets (url) VALUES ('basement')")
	if createErr != nil {
		return
	}

}
