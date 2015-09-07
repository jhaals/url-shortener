package main

import (
	"math"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
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

func index(response http.ResponseWriter, request *http.Request) {

	http.Error(response, "fail", 400)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	http.Handle("/", r)
	http.ListenAndServe(":1337", handlers.LoggingHandler(os.Stdout, r))
}
