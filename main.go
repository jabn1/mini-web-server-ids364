package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

var messages map[int]string

func main() {
	port := "5000"
	messages = map[int]string{
		1: "Hola",
		2: "Mundo",
	}
	r := registerRoutes()
	fmt.Println("Listening on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func registerRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/messages", getMessages)
		r.Post("/messages", createMessage)
		r.Put("/messages/{msgId}", updateMessage)
		r.Delete("/messages/{msgId}", deleteMessage)
	})
	return r
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(messages)
}

func createMessage(w http.ResponseWriter, r *http.Request) {

}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	msgId := chi.URLParam(r, "msgId")
	body, err := ioutil.ReadAll(r.Body)
	if msgId == "" {
		http.Error(w, "Invalid message id", 400)
		return
	}
	if err != nil {
		http.Error(w, "Invalid message", 400)
	}
	if len(body) == 0 {
		http.Error(w, "Empty request body", 400)
	}
	id, err := strconv.Atoi(msgId)
	if err != nil {
		http.Error(w, "Invalid message id", 400)
	}
	messages[id] = string(body)
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {

}
