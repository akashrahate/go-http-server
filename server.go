package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type welcome string

func (c welcome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to our server!")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login")
}
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Log out success")
}
func getJson(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "getJson called")

	switch r.Method {
	case "GET":
		w.Write([]byte(`"message": "GET is called"`))
	case "POST":
		w.Write([]byte(`"message": "POST is called!"`))
	}
}

func checkUser(w http.ResponseWriter, r *http.Request) {

	var check user
	dbpassword := "Pass@123"

	err := json.NewDecoder(r.Body).Decode(&check)
	if err != nil {
		log.Fatal("error decoding into struct")
	}
	if check.Password == dbpassword {
		fmt.Fprintf(w, "success logging in")
		log.Printf("passwword matched")
	}
	if check.Password != dbpassword {
		fmt.Fprintf(w, "unsuccessful attempt")
		log.Printf("wrong password")
	}
}

func main() {
	//Router
	router := http.NewServeMux()

	//Handler
	var wc welcome
	router.Handle("/", wc)

	//Handler Funcs
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/login", login)
	router.HandleFunc("/json", getJson)
	router.HandleFunc("/check", checkUser)
	//Server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//Run Server
	fmt.Printf("server starting %s", server.Addr)
	server.ListenAndServe()

}
