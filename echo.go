package main

import (
	"log"
	"net/http"
	"strings"
)

// Service defines a service.
type Service func(string) string

// Handler is an HTTP Handler.
type Handler struct {
	service Service
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	segments := strings.Split(r.URL.Path, "/")
	_, err := w.Write([]byte(h.service(segments[len(segments)-1])))

	if err != nil {
		log.Println(err)
	}
}

// NewHandler creates a new Handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func echoService(s string) string {
	return strings.ToUpper(s)
}

func reverseService(s string) string {
	result := ""

	for _, v := range s {
		result = string(v) + result
	}

	return result
}

func main() {
	echoHandler := NewHandler(echoService)
	reverseHandler := NewHandler(reverseService)

	http.Handle("/echo/", echoHandler)
	http.Handle("/reverse/", reverseHandler)
	port := ":8080"

	log.Println("starting server at port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
