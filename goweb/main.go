package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

func HandleMain(wr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wr, req.URL.Path)
}

func TodoIndex(wr http.ResponseWriter, req *http.Request) {

	todoId := req.FormValue("todoId")

	if todoId == "" { // print it
		todos := Todos{
			Todo{Name: "Write presentation"},
			Todo{Name: "Host meetup"},
		}

		wr.Header().Set("Content-Type", "application/json;charset=UTF-8")
		wr.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(wr).Encode(todos); err != nil {
			panic(err)
		}
	} else { // grab it.
		fmt.Fprintln(wr, "Todo ID: "+todoId)
	}

}

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)

		log.Printf("%s\t%s\t%s",
			r.Method,
			r.RemoteAddr,
			r.RequestURI)
	})
}

func main() {
	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/todos", TodoIndex)

	fmt.Println("Running...")
	http.ListenAndServe(":6455", Logger(http.DefaultServeMux))
}
