package server

import (
	"fmt"
	"net/http"
)

func Middleware1(next http.Handler) http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {
		//before
		fmt.Println("Before ..")
		next.ServeHTTP(w, r)
		//after
		fmt.Println("After ..")
	})
}
