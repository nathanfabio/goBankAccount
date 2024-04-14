package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	addr string
}

func NewAddress(addr string) *API {
	return &API{addr: addr}
}

// func WJson(w http.ResponseWriter, status int, v any) error {
// 	w.WriteHeader(status)
// 	w.Header().Set("Content-Type", "application/json")
// 	return json.NewEncoder(w).Encode(v)
// }

// type apiFunc func(http.ResponseWriter, *http.Request) error

// type ApiError struct {
// 	Error string
// }

// func HttpHandleFunc(a apiFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := a(w, r); err != nil {
// 			WJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
// 		}
// 	}
// }

// func toGinHandler(h http.HandlerFunc) gin.HandlerFunc {
// 	return func (c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

func (a *API) Run() {
	router:= gin.Default()
	router.Handle("GET", "/account", a.handleGetAccount)
	// router.Handle("GET", "/account/{id}", toGinHandler(HttpHandleFunc(a.handleGetAccount)))
	
	log.Println("Listening on", a.addr)

	router.Run(a.addr)
}

// func (a *API) handleAccount(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

func (a *API) handleGetAccount(c *gin.Context){
	// id:= c.Param("id")

	account:= NewAccount("Nathan", "Fabio")

	c.IndentedJSON(http.StatusOK, account)
}

func (a *API) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
