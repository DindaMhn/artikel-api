package repositories

import "artikel/master/model"

type ArtikelRepo interface {
	GetAllDataArtikel() ([]*model.ArtikelModel, error)
	GetArtikelById(id string) (model.ArtikelModel, error)
	UpdateArtikelById(id string, DataArtikel model.ArtikelModel) error
	DeleteArtikelById(id string) error
	InsertArtikel(DataArtikel model.ArtikelModel) error
}
