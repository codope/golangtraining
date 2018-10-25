package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)                       // with function
	r.Handle("/panic", recoverMiddleware(panicHandler())) // with Handler

	port := ":9090"
	fmt.Printf("Open the browser at the root page: localhost%s/\n", port)
	fmt.Printf("And then with the apihandler: localhost%s/panic\n", port)
	http.ListenAndServe(port, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("in indexHandler")
}

func panicHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("forcing a panic")
	})
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer the recovery
		defer func() {
			if r := recover(); r != nil {
				log.Printf("There was a panic with this message: %v\n", r)
				log.Printf("But we recovered from it in the middleware and so the server is still up.\n")
				switch x := r.(type) {
				case string:
					http.Error(w, r.(string), http.StatusInternalServerError)
				case error:
					err := x
					http.Error(w, err.Error(), http.StatusInternalServerError)
				default:
					http.Error(w, "unknown panic", http.StatusInternalServerError)
				}
			} else {
				log.Printf("There was no panic.\n")
			}
		}()

		// call the actual api handler
		next.ServeHTTP(w, r)
	})
}

func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
