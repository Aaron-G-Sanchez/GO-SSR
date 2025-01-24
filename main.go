package main

import (
	"fmt"
	"log"
	"net/http"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := Page(global.Count)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if r.Form.Has("global") {
		global.Count++
	}

	getHandler(w, r)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}

		getHandler(w, r)
	})

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Printf("Error listening: %v", err)
	}
}
