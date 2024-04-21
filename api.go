package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	addr string
	store Storage
}

func NewAddress(addr string, store Storage) *API {
	return &API{addr: addr,
		store: store}
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
	// router.Handle("GET", "/account", a.handleGetAccount)
	router.Handle("GET", "/account/{id}", (a.handleGetAccountByID))
	router.Handle("GET", "/accounts", (a.handleGetAccount))
	router.Handle("POST", "/account", (a.handleCreateAccount))
	
	log.Println("Listening on", a.addr)

	router.Run(a.addr)
}

var request Account
// func (a *API) handleAccount(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
func (a *API) handleGetAccount(c *gin.Context) {
	accounts, err := a.store.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, accounts)
}

func (a *API) handleGetAccountByID(c *gin.Context){
	id:= c.Param("id")
	fmt.Println(id)

	c.ShouldBindJSON(&request)
	c.JSON(http.StatusOK, &request)
}

func (a *API) handleCreateAccount(c *gin.Context) {
	createAccountReq:= AccountRequest{}

	if err:= json.NewDecoder(c.Request.Body).Decode(&createAccountReq); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account:= NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
	if err := a.store.CreateAccount(account); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, account)
}

func (a *API) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
