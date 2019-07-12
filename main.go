package main

import (
	"DSLoginServer/app"
	"DSLoginServer/controllers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var router *mux.Router
	router = mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	http.Handle("/", router)

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST","OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With",  "Access-Control-Allow-Headers", "Authorization"}),
	)(router)


	log.Fatal(http.ListenAndServe(":8082", handler))
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