package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type data struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", HelloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	file, _ := ioutil.ReadFile("./storage/data.json")
	storage := data{}

	_ = json.Unmarshal([]byte(file), &storage)

	fmt.Fprintf(w, "Hello, my name is %s!", storage.Name)

}
