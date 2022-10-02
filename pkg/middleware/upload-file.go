package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("image")

		if err != nil && r.Method == "PATCH" {
			ctx := context.WithValue(r.Context(), "dataFile", "false")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("error retrieving the file")
			return
		}
		defer file.Close()

		const MAX_UPLOAD_SIZE = 10 << 20
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "file cannot above 10MB"}
			json.NewEncoder(w).Encode(response)
			return
		}

		tempFile, err := ioutil.TempFile("uploads", "file-*-"+strings.ReplaceAll(handler.Filename, " ", "_"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:] // split uploads/

		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
