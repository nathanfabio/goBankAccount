package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func (a *API) Run() {
	router:= gin.Default()
	router.Handle("GET", "/accounts", (a.handleGetAccount))
	router.Handle("GET", "/account/:id", (a.handleGetAccountByID))
	router.Handle("POST", "/account", (a.handleCreateAccount))
	router.Handle("DELETE", "/account/:id", (a.handleDeleteAccount))
	router.Handle("GET", "/transfer", (a.handleTransfer))
	
	log.Println("Listening on", a.addr)

	router.Run(a.addr)
}

func (a *API) handleGetAccount(c *gin.Context) {
	accounts, err := a.store.GetAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, accounts)
}

func (a *API) handleGetAccountByID(c *gin.Context){
	id:= c.Param("id")
	clientID, err:= strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	account, err := a.store.GetAccountByID(clientID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	c.JSON(http.StatusOK, account)
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

func (a *API) handleDeleteAccount(c *gin.Context) {
	id:= c.Param("id")
	clientID, err:= strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	if err:= a.store.DeleteAccount(clientID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting account"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "account deleted"})
}

func (a *API) handleTransfer(c *gin.Context) {
	transferReq:= Tranfer{}
	if err:= json.NewDecoder(c.Request.Body).Decode(&transferReq); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer c.Request.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "transfer done"})
}
