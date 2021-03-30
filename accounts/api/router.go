package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	accountsApi *AccountsApi
}

func InitRouter(accountsApi *AccountsApi) *Router {
	return &Router{
		accountsApi: accountsApi,
	}
}

func (r *Router) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	router.HandleFunc("/signup", r.accountsApi.SignUp).Methods("POST")

	return router
}
