package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/Prakhyath07/GoPracticeProjects/bookStoreManagement/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func main(){
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)

	fmt.Println("Starting server at port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}