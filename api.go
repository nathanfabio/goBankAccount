package main

import "net/http"

type API struct {
	addr string
}

func NewAddress(addr string) *API {
	return &API{addr: addr}
}

func (a *API) handleAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (a *API) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
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
