package router

import (
	"github.com/Flyingmonk01/go-pro/internal/controller"
	"github.com/gorilla/mux"
)

func AuthRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/auth", controller.AuthHandler).Methods("GET")
	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	return router
}
