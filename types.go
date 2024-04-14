package main

import "math/rand"

type Account struct {
	ID         int `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	NumberBank int64  `json:"numberBank"`
	Balance    float64 `json:"balance"`
}

// NewAccount creates a new bank account
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		FirstName: firstName,
		LastName:  lastName,
		NumberBank: rand.Int63n(10000),
		Balance:   0,
	}
}