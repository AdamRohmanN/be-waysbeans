package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"waysbeans/dto"
	profiledto "waysbeans/dto/profile"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerProfile struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *handlerProfile {
	return &handlerProfile{ProfileRepository}
}

func convertResponseProfile(u models.Profile) profiledto.ProfileResponse {
	return profiledto.ProfileResponse{
		Image:    u.Image,
		Address:  u.Address,
		Postcode: u.Postcode,
		User:     u.User,
	}
}

func (h *handlerProfile) GetProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var profile models.Profile
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProfile(profile)}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerProfile) CreateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	req := profiledto.ProfileRequest{
		Address:  r.FormValue("address"),
		Postcode: r.FormValue("postcode"),
	}

	validation := validator.New()
	err := validation.Struct(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	profile := models.Profile{
		Image:    filename,
		Address:  req.Address,
		Postcode: req.Postcode,
		UserId:   userId,
	}

	profile, err = h.ProfileRepository.CreateProfile(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	profile, _ = h.ProfileRepository.GetProfile(profile.Id)
	profile.Image = os.Getenv("PATH_IMAGE") + profile.Image

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: profile}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerProfile) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	req := profiledto.ProfileRequest{
		Image:    filename,
		Address:  r.FormValue("address"),
		Postcode: r.FormValue("postcode"),
	}

	validation := validator.New()
	err := validation.Struct(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	profile, _ := h.ProfileRepository.GetProfile(id)

	if filename != "false" {
		profile.Image = filename
	}
	if req.Address != "" {
		profile.Address = req.Address
	}
	if req.Postcode != "" {
		profile.Postcode = req.Postcode
	}

	profile, err = h.ProfileRepository.UpdateProfile(profile, profile.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: profile}
	json.NewEncoder(w).Encode(res)
}
