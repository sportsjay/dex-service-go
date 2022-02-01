package handlers

import (
	"encoding/json"
	"net/http"
)

func HOrderGet(w http.ResponseWriter, r *http.Request) {
	data := "Get Orders"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func HOrderPost(w http.ResponseWriter, r *http.Request) {
	/**
	Order scheme: order limit
	token A to be traded into token B
	n*[A] => m*[B]
	*/

	// check if account exists

	// do look up for token exchange contract

	// check if token A suffice

	// check if token B is in "vault"

	// if amount B available in "vault", assign token B ordered by time and amount
	// note: amount must be equal

	// if n > 0 after transaction, save to order table
}
