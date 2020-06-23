package main

import (
	"DSAppServer/controllers"
	"DSAppServer/dbh"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	//var router *mux.Router
	router := *chi.NewRouter()
	//router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	dbh.InitDB()

	//router = mux.NewRouter()
	//router.Use(app.JwtAuthentication)


	c := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With",  "Access-Control-Allow-Headers"},
		// AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(c.Handler)

	router.Post("/api/users/new", controllers.CreateUser)
	router.Post("/api/users/login", controllers.Authenticate)
	router.Post("/api/host/app", controllers.DomainApp)
	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+ port, &router))
}



func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only work if you are accessing the server from your local machine.
	print("dickballs")
	http.Redirect(w, r, "https://192.168.0.2:8081"+r.RequestURI, http.StatusMovedPermanently)
}

func testTimeString() {
	t1, err := time.Parse(
		time.RFC3339,
		"2018-03-13T00:17:41+00:00")
	fmt.Print(t1)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
}