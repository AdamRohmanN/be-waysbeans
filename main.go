package main

import (
	"fmt"
	"net/http"
	"os"
	"waysbeans/database"
	"waysbeans/pkg/mysql"
	"waysbeans/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load .env file")
	}

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running on " + os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("PORT"), r)
}
