package main

import (
		"./data"
		"net/http"
		"errors"
	)

func session(w http.ResponseWritter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if 0k, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}