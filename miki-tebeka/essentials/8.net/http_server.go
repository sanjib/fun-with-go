package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type DBRecord struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type DBData struct {
	mu   sync.Mutex
	data map[string]interface{}
}

var dbData = DBData{
	data: make(map[string]interface{}),
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func db(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dbPost(w, r)
	} else if r.Method == "GET" {
		dbGet(w, r)
	}
}

func dbGet(w http.ResponseWriter, r *http.Request) {
	pathBits := strings.Split(r.URL.Path, "/")
	key := pathBits[len(pathBits)-1]

	dbData.mu.Lock()
	defer dbData.mu.Unlock()

	val, _ := dbData.data[key]
	fmt.Fprintf(w, "key:%v;val:%v", key, val)
}

// Example POST body
// {"key":"x", "value":1}
func dbPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var dbRecord DBRecord
	err := decoder.Decode(&dbRecord)
	if err != nil {
		log.Println(err)
	}

	dbData.mu.Lock()
	defer dbData.mu.Unlock()
	dbData.data[dbRecord.Key] = dbRecord.Value

	fmt.Fprintf(w, "%+v", dbRecord)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/db/", db)
	srv := http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	log.Println("Starting server at :4000...")
	err := srv.ListenAndServe()
	log.Println(err)
}
