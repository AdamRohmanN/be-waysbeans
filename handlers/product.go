package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"waysbeans/dto"
	productdto "waysbeans/dto/product"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func convertResponseProduct(u models.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		Image:    u.Image,
		Name:     u.Name,
		Desc:     u.Desc,
		Price:    u.Price,
		Stock:    u.Stock,
		User:     u.User,
		Category: u.Category,
	}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	for i, p := range products {
		products[i].Image = os.Getenv("PATH_IMAGE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	product.Image = os.Getenv("PATH_IMAGE") + product.Image

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(product)}
	json.NewEncoder(w).Encode(res)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// upload image
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	stock, _ := strconv.Atoi(r.FormValue("stock"))
	category_id, _ := strconv.Atoi(r.FormValue("category_id"))

	req := productdto.ProductRequest{
		Name:       r.FormValue("name"),
		Desc:       r.FormValue("desc"),
		Price:      price,
		Stock:      stock,
		CategoryId: category_id,
	}

	validation := validator.New()
	err := validation.Struct(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	product := models.Product{
		Image:  filename,
		Name:   req.Name,
		Desc:   req.Desc,
		Price:  req.Price,
		Stock:  req.Stock,
		UserId: userId,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(res)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.Id)

	product.Image = os.Getenv("PATH_IMAGE") + product.Image

	w.WriteHeader(http.StatusOK)
	res := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(res)
}
