package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zyzmoz/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBooksRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), r))
	log.Println("Server listening on " + os.Getenv("APP_PORT"))
}
