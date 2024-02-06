package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	urlMap = map[string]string{}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var urlMap map[string]string

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		createShortURL(w, r)
	} else if r.Method == http.MethodGet {
		getOriginalURL(w, r)
	}
}

func createShortURL(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	r.Body.Close()
	if len(body) == 0 {
		http.Error(w, "Body is empty.", http.StatusBadRequest)
		return
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	shortURL := randString(10)
	urlMap[shortURL] = string(body)

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v://%v%v%v", scheme, r.Host, r.RequestURI, shortURL)
}

func getOriginalURL(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	shortURL := strings.Split(path, "/")[0]

	if fullURL, ok := urlMap[shortURL]; ok {
		w.Header().Set("Location", fullURL)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
