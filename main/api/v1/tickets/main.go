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

func init() {

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

	db, err := sql.Open("mysql", "goexp:F76MWx3Px2@/exp_golang")
	if err != nil {
		return
	}

	_, createErr := db.Query("CREATE TABLE `tickets` (`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY, `url` varchar(50) NOT NULL)")
	if createErr != nil {
		return
	}

}
