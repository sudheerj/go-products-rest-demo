package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sudheerj/go-rest.git/model"
	"net/http"
	"strconv"

	"github.com/sudheerj/go-rest.git/storage"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	if limit < 1 {
		limit = 10
	}

	orders, err := storage.DBStore.GetProducts(limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, orders)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := model.Product{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if err := storage.DBStore.CreateProduct(&product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, product)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	product, fetchErr := storage.DBStore.GetProduct(name);
	if fetchErr != nil {
		respondWithError(w, http.StatusInternalServerError, fetchErr.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	product := model.Product{}
	decoder := json.NewDecoder(r.Body)
	if updateErr := decoder.Decode(&product); updateErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	product.Name = name

	if err := storage.DBStore.UpdateProduct(&product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if err := storage.DBStore.DeleteProduct(name); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
