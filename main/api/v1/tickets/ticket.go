package tickets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func init() {

}

func GetList(w http.ResponseWriter, r *http.Request) {

	file, _ := ioutil.ReadFile("./storage/data.json")
	storage := data{}

	_ = json.Unmarshal([]byte(file), &storage)

	_, err := fmt.Fprintf(w, "Hello, my name is %s %s!", storage.Name, storage.LastName)
	if err != nil {
		return
	}

}
