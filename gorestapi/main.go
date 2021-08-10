package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gorestapi/models"
	"gorestapi/utils"
	"gorestapi/warehouse"
)

func main() {

	http.HandleFunc("/products", GetProduct)
	http.HandleFunc("/products/create", PostProduct)
	http.HandleFunc("/products/update", UpdateProduct)
	http.HandleFunc("/products/delete", DeleteProduct)
	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// Get Product Terbaru
func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		sortBy := r.URL.Query().Get("sortBy")

		products, err := warehouse.GetAll(ctx, sortBy)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, products, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

// Post Product
func PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var prdct models.Product

		if err := json.NewDecoder(r.Body).Decode(&prdct); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := warehouse.Insert(ctx, prdct); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Put Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var prdct models.Product

		if err := json.NewDecoder(r.Body).Decode(&prdct); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		fmt.Println(prdct)

		if err := warehouse.Update(ctx, prdct); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

// Delete Product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var prdct models.Product

		product_id := r.URL.Query().Get("product_id")

		if product_id == "" {
			utils.ResponseJSON(w, "product_id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		prdct.ProductId, _ = strconv.Atoi(product_id)

		if err := warehouse.Delete(ctx, prdct); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
