package main

import (
	"fmt"
	"net/http"
	"waysbeans/database"
	"waysbeans/pkg/mysql"
	"waysbeans/routes"

	"github.com/gorilla/mux"
)

func main() {
	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	port := "localhost:5000"
	fmt.Println("server running on " + port)
	http.ListenAndServe(port, r)
}
