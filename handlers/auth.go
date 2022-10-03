package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"waysbeans/dto"
	authdto "waysbeans/dto/auth"
	"waysbeans/models"
	"waysbeans/pkg/bcrypt"
	jwtToken "waysbeans/pkg/jwt"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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
		Status:   "customer",
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
	res := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	//check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	//check password
	isValid := bcrypt.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		res := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(res)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("unauthorized")
		return
	}

	loginResponse := authdto.LoginResponse{
		Status:   user.Status,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	user, err := h.AuthRepository.GetAuth(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		Id:     user.Id,
		Status: user.Status,
		Name:   user.Name,
		Email:  user.Email,
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(res)
}
