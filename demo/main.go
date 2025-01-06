package main

import (
	"net/http"

	"github.com/haitux/clam"
)

var authService *clam.AuthService

func main() {
	authService = clam.NewAuthService("pandora", "io.haitu.pandora", "KMTM6jyCtBI6YnsVvHGtZfo3ACxatEfL")
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":9000", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	authService.Login(w, r)
}
