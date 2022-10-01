package handlers

import (
	"encoding/json"
	"net/http"
	"waysbeans/dto"
	authdto "waysbeans/dto/auth"
	"waysbeans/models"
	"waysbeans/pkg/bcrypt"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	validation := validator.New()
	err := validation.Struct(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	password, err := bcrypt.HashingPassword(req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseUser(data)}
	json.NewEncoder(w).Encode(res)
}
