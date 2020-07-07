package controller

import (
	"artikel/master/model"
	"artikel/master/tools"
	"artikel/master/usecase"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var ArtikelResponse tools.Response

type ArtikelHandler struct {
	ArtikelUseCase usecase.ArtikelUseCase
}

func ArtikelController(r *mux.Router, service usecase.ArtikelUseCase) {
	ArtikelHandler := ArtikelHandler{service}
	r.HandleFunc("/Artikel", ArtikelHandler.ListArtikel).Methods(http.MethodGet)
	r.HandleFunc("/Artikel/{id}", ArtikelHandler.ArtikelById).Methods(http.MethodGet)
	r.HandleFunc("/Artikel/{id}", ArtikelHandler.DeleteById).Methods(http.MethodDelete)
	r.HandleFunc("/Artikel/{id}", ArtikelHandler.UpdateArtikel).Methods(http.MethodPut)
	r.HandleFunc("/Artikel", ArtikelHandler.InsertArtikel).Methods(http.MethodPost)
}

func (s ArtikelHandler) ListArtikel(w http.ResponseWriter, r *http.Request) {
	Artikel, err := s.ArtikelUseCase.GetAllDataArtikel()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	} else {
		ArtikelResponse = tools.Response{"success", Artikel}
		byteOfArtikels, err := json.Marshal(ArtikelResponse)
		if err != nil {
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfArtikels)
	}

}
func (s ArtikelHandler) ArtikelById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idArtikel := param["id"]
	Artikel, err := s.ArtikelUseCase.GetArtikelById(idArtikel)
	if err != nil {
		if err == sql.ErrNoRows {
			ArtikelResponse.Message = "Id Not Found"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ArtikelResponse.Message))
		}
	} else {
		ArtikelResponse = tools.Response{"success", Artikel}
		byteOfArtikels, err := json.Marshal(ArtikelResponse)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Oops something when wrong"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfArtikels)
	}
}

func (s ArtikelHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	var Artikel model.ArtikelModel
	param := mux.Vars(r)
	idArtikel := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&Artikel)
	err := s.ArtikelUseCase.DeleteArtikelById(idArtikel)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Deleted"))
	}

}
func (s ArtikelHandler) UpdateArtikel(w http.ResponseWriter, r *http.Request) {
	var Artikel model.ArtikelModel
	param := mux.Vars(r)
	idArtikel := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&Artikel)
	err := s.ArtikelUseCase.UpdateArtikelById(idArtikel, Artikel)
	if err != nil {
		w.Write([]byte("Id Not Found"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data Updated"))
	}
}
func (s ArtikelHandler) InsertArtikel(w http.ResponseWriter, r *http.Request) {
	var Artikel model.ArtikelModel
	_ = json.NewDecoder(r.Body).Decode(&Artikel)
	err := s.ArtikelUseCase.InsertArtikel(Artikel)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Insert Failed"))

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Insert Success"))
	}
}
