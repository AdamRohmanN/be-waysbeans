package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
	r.HandleFunc("/profile", middleware.Auth(middleware.UploadFile(h.CreateProfile))).Methods("POST")
	r.HandleFunc("/profile/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProfile))).Methods("PATCH")
}
