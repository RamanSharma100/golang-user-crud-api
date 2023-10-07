package main

import (
	"crudapi/pkgs/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.UserRoutes(r)

	http.Handle("/", r)

	log.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
