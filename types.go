package main

import (
	"math/rand"
	"time"
)

type Tranfer struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type AccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	NumberBank int64     `json:"numberBank"`
	Balance    float64   `json:"balance"`
	CreateAt   time.Time `json:"createAt"`
}

// NewAccount creates a new bank account
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName:  firstName,
		LastName:   lastName,
		NumberBank: rand.Int63n(10000),
		Balance:    0,
		CreateAt:   time.Now().Local().UTC(),
	}
}
