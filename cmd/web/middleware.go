package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit The Page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf is the CSRF protection middleware

func NoSurf(next http.Handler) http.Handler {

	//	fmt.Println("Inside NoSurf")
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

// LoadSession loads and saves session data for request.
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
