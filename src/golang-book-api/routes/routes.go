package routes

import (
	con "golang-book-api/controllers"

	"github.com/gorilla/mux"
)

// [GET] - api/v1/books
func routeGetBooks(router *mux.Router) {
	router.HandleFunc("/api/v1/books", con.GetBooks).Methods("GET")
}

// [GET] - api/v1/books/{id} [params]=[ID]
func routeGetBook(router *mux.Router) {
	router.HandleFunc("/api/v1/books/{id}", con.GetBook).Methods("GET")
}

// [POST] - api/v1/books/add [body]=[Book object]
func routeCreateBook(router *mux.Router) {
	router.HandleFunc("/api/v1/books/add", con.CreateBook).Methods("POST")
}

// [PUT] - api/v1/books/{id} [params]=[ID]
func routeUpdateBook(router *mux.Router) {
	router.HandleFunc("/api/v1/books/{id}", con.UpdateBook).Methods("PUT")
}

// [DELETE] - api/v1/books/{id} [params]=[ID]
func routeDeleteBook(router *mux.Router) {
	router.HandleFunc("/api/v1/books/{id}", con.DeleteBook).Methods("DELETE")
}

// SetRoutes sets all API routes
func SetRoutes(router *mux.Router) {
	routeGetBooks(router)
	routeGetBook(router)
	routeCreateBook(router)
	routeUpdateBook(router)
	routeDeleteBook(router)
}
