package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var task int

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %d", task)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type PostTaskHandlerRequest struct {
	Task string `json:"task"`
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req PostTaskHandlerRequest
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newTask, err := strconv.Atoi(req.Task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	task = newTask
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /task", http.HandlerFunc(GetTaskHandler))
	mux.Handle("POST /task", http.HandlerFunc(PostTaskHandler))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
