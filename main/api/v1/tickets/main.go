package tickets

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
func (db *dbConn) connect(w http.ResponseWriter) {
	conn, err := sql.Open("mysql", db.user+":"+db.pass+"@tcp(mysql_storage)/"+db.base)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	db.conn = conn
}

func getConnection(w http.ResponseWriter) *sql.DB {
	db := dbConn{
		user: "goexp",
		pass: "F76MWx3Px2",
		base: "exp_golang",
	}
	db.connect(w)
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

	db := getConnection(w)

	tickets, err := db.Query("SELECT * FROM tickets")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	for tickets.Next() {
		var (
			id  int
			url string
		)
		if err := tickets.Scan(&id, &url); err != nil {
			log.Fatal(err)
		}
		log.Println(id, url)
	}

}

func Create(w http.ResponseWriter, r *http.Request) {

	db := getConnection(w)

	_, err := db.Query("CREATE TABLE `tickets` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `url` varchar(50) NOT NULL)")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Creating")

}

func Add(w http.ResponseWriter, r *http.Request) {

	db := getConnection(w)
	_, err := db.Query("INSERT INTO tickets (url) VALUES ('basement')")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, "Added")

}
