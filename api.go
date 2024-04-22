package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
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
	router.GET("/account/:id", withJWTAuth, a.handleGetAccountByID)
	router.Handle("POST", "/account", (a.handleCreateAccount))
	router.Handle("DELETE", "/account/:id", (a.handleDeleteAccount))
	router.Handle("GET", "/transfer", (a.handleTransfer))
	
	log.Println("Listening on", a.addr)

	router.Run(a.addr)
}

// const JWTKey = "golangjwt" //need to stay out of envaronment

func withJWTAuth(c *gin.Context) {
	tokenString:= c.GetHeader("Authorization")
	token, err:= validateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Set("token", token)
	c.Next()
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret:= os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func createJWT(account *Account) (string, error) {
	secret:= os.Getenv("JWT_SECRET")
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accountNumber": account.NumberBank,
		"expire": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func (a *API) handleGetAccount(c *gin.Context) {
	accounts, err := a.store.GetAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, accounts)
}

func (a *API) handleGetAccountByID(c *gin.Context){
	token, exists := c.Get("token")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found or invalid"})
        return
    }
	c.Set("token", token)
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

	token, err:= createJWT(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("token: ", token)

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