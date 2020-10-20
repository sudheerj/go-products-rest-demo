package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sudheerj/go-rest.git/model"
	"net/http"
	"strconv"

	"github.com/sudheerj/go-rest.git/storage"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["title"]
	product, prodErr := storage.DBStore.GetProduct(productName)
	if prodErr != nil {
		respondWithError(w, http.StatusInternalServerError, prodErr.Error())
		return
	}

	reviews, err := storage.DBStore.GetReviews(product);
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, reviews)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productName := vars["title"]
	product, prodErr := storage.DBStore.GetProduct(productName)
	if prodErr != nil {
		respondWithError(w, http.StatusInternalServerError, prodErr.Error())
		return
	}

	review := model.Review{ProductID: product.ID}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&review); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := storage.DBStore.CreateReview(&review); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, review)
}

func GetReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	review, fetchErr := storage.DBStore.GetReview(id);
	if fetchErr != nil {
		respondWithError(w, http.StatusInternalServerError, fetchErr.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, review)
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productName, _ := vars["name"]
	_, prodErr := storage.DBStore.GetProduct(productName)
	if prodErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product name")
		return
	}
	id, _ := strconv.Atoi(vars["id"])
	review, getErr := storage.DBStore.GetReview(id)
	if getErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid review id")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if updateErr := decoder.Decode(&review); updateErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := storage.DBStore.UpdateReview(review); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productName, _ := vars["name"]
	_, prodErr := storage.DBStore.GetProduct(productName)
	if prodErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product name")
		return
	}
	id, _ := strconv.Atoi(vars["id"])
	review, getErr := storage.DBStore.GetReview(id)
	if getErr != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid review id")
		return
	}

	if err := storage.DBStore.DeleteReview(review); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
