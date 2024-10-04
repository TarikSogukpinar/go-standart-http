package router

import (
	"database/sql"
	"go-standart-http/internal/handlers"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", handlers.GetUsersHandler(db))
	mux.HandleFunc("/users/create", handlers.CreateUserHandler(db))
	mux.HandleFunc("/users/update", handlers.UpdateUserHandler(db))
	mux.HandleFunc("/users/delete", handlers.DeleteUserHandler(db))

	return mux
}
