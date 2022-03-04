package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func(){
			if err := recover(); err != nil {
				rw.Header().Set("Connection", "close")
				app.serverErrorResponse(rw, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(rw, r)
	})
}