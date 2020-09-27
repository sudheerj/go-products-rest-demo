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
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if limit < 1 {
		limit = 10
	}

	orders, err := storage.DBStore.GetProducts(limit, start)
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
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	product, fetchErr := storage.DBStore.GetProduct(id);
	if fetchErr != nil {
		respondWithError(w, http.StatusInternalServerError, fetchErr.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}
	product := model.Product{}
	decoder := json.NewDecoder(r.Body)
	if updateErr := decoder.Decode(&product); updateErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	product.ID = id

	if err := storage.DBStore.UpdateProduct(&product); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := storage.DBStore.DeleteProduct(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
