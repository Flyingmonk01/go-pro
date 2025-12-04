package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func UserRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/user", UserHandler).Methods("GET")
	return router
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ok"}`))
}
