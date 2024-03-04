package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-5-pbi-fullstack-developer-adityap/helpers"
	"task-5-pbi-fullstack-developer-adityap/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helpers.ResponseJson
var ResponseError = helpers.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var photos []models.Photo

	if err := models.DB.Find(&photos).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, photos)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var photo models.Photo
	if err := models.DB.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Data Tidak Ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, photo)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photo); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&photo).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, photo)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var photo models.Photo

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photo); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&photo).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak Dapat Update Data")
		return
	}

	photo.Id = id
	ResponseJson(w, http.StatusOK, photo)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var photo models.Photo
	if models.DB.Delete(&photo, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak Dapat Menghapus Data")
		return
	}

	response := map[string]string{"message": "Data Berhasil Dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
