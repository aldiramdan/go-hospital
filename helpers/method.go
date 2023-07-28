package helpers

import (
	"errors"
	"net/http"
)

func MethodGet(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errors.New("method not allowed")
	}
	return nil
}

func MethodPost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.New("method not allowed")
	}
	return nil
}

func MethodPut(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return errors.New("method not allowed")
	}
	return nil
}

func MethodDelete(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodDelete {
		return errors.New("method not allowed")
	}
	return nil
}
