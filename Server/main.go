package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	users = []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
)

func main() {

	http.HandleFunc("/users", authMiddleware(loggerMiddleware(handleUsers)))

	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}

}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			log.Printf("[%s] %s - error: userID is not provided", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", userID)

		r = r.WithContext(ctx)

		next(w, r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromCx := r.Context().Value("userID")
		userID, ok := idFromCx.(string)
		if !ok {
			log.Printf("[%s] %s - error: userID is invalid", r.Method, r.URL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("%s [%s] %s by userID %s\n", r.RemoteAddr, r.Method, r.URL, userID)
		next(w, r)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var user User

	if err := json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	users = append(users, user)
}
