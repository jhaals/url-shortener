package main

import (
	"math"
	"strings"
	"net/http"
	"os"
	"encoding/json"
	"log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

func encodeHandler(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data struct {
		URL    string `json:"url"`
	}
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(response, `{"error": "Unable to parse json"}`, http.StatusBadRequest)
		return
	}

	resp := map[string]string{"url": data.URL, "error": ""}
	jsonData, _ := json.Marshal(resp)
	response.Write(jsonData)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/encode", encodeHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":1337", handlers.LoggingHandler(os.Stdout, r)))
}
