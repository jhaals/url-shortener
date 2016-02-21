package main

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"os"
	"path"
	"strings"
)

const base string = "0123456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

func encode(n int64) string {
	var s string
	var num = float64(n)

	for num > 0 {
		s = string(base[int(num)%len(base)]) + s
		num = math.Floor(num / float64(len(base)))
	}

	return s
}

func decode(s string) int {
	var num = 0
	for _, element := range strings.Split(s, "") {
		num = num*len(base) + strings.Index(base, element)
	}

	return num
}

func decodeHandler(response http.ResponseWriter, request *http.Request, db Database) {
	id := decode(mux.Vars(request)["id"])
	url, err := db.Get(id)
	if err != nil {
		http.Error(response, `{"error": "No such URL"}`, http.StatusNotFound)
		return
	}
	http.Redirect(response, request, url, 301)
}

func encodeHandler(response http.ResponseWriter, request *http.Request, db Database, baseURL string) {
	decoder := json.NewDecoder(request.Body)
	var data struct {
		URL string `json:"url"`
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(response, `{"error": "Unable to parse json"}`, http.StatusBadRequest)
		return
	}

	if !govalidator.IsURL(data.URL) {
		http.Error(response, `{"error": "Not a valid URL"}`, http.StatusBadRequest)
		return
	}

	id, err := db.Save(data.URL)
	if err != nil {
		log.Println(err)
		return
	}

	resp := map[string]string{"url": baseURL + encode(id), "id": encode(id), "error": ""}
	jsonData, _ := json.Marshal(resp)
	response.Write(jsonData)

}

func main() {

	if os.Getenv("BASE_URL") == "" {
		log.Fatal("BASE_URL environment variable must be set")
	}
	if os.Getenv("DB_PATH") == "" {
		log.Fatal("DB_PATH environment variable must be set")
	}
	db := sqlite{Path: path.Join(os.Getenv("DB_PATH"), "db.sqlite")}
	db.Init()

	baseURL := os.Getenv("BASE_URL")
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	r := mux.NewRouter()
	r.HandleFunc("/save",
		func(response http.ResponseWriter, request *http.Request) {
			encodeHandler(response, request, db, baseURL)
		}).Methods("POST")
	r.HandleFunc("/{id}", func(response http.ResponseWriter, request *http.Request) {
		decodeHandler(response, request, db)
	})
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	log.Println("Starting server on port :1337")
	log.Fatal(http.ListenAndServe(":1337", handlers.LoggingHandler(os.Stdout, r)))
}
