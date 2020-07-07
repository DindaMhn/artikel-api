package usecase

import (
	"artikel/master/model"
	"artikel/master/repositories"
	"artikel/master/tools"
)

type ArtikelUsecaseImpl struct {
	ArtikelRepo repositories.ArtikelRepo
}

func (au ArtikelUsecaseImpl) GetAllDataArtikel() ([]*model.ArtikelModel, error) {
	Artikel, err := au.ArtikelRepo.GetAllDataArtikel()
	if err != nil {
		return nil, err
	}
	return Artikel, nil
}
func (s ArtikelUsecaseImpl) GetArtikelById(id string) (model.ArtikelModel, error) {
	Artikel, err := s.ArtikelRepo.GetArtikelById(id)

	return Artikel, err
}
func (s ArtikelUsecaseImpl) UpdateArtikelById(id string, DataArtikel model.ArtikelModel) error {
	err := tools.ValidationNotNil(DataArtikel.IdUser, DataArtikel.IdCategory, DataArtikel.Description,
		DataArtikel.ArtikelTitle, DataArtikel.Date)
	if err != nil {
		return err
	}
	err = s.ArtikelRepo.UpdateArtikelById(id, DataArtikel)
	if err != nil {
		return err
	}
	return err
}
func (s ArtikelUsecaseImpl) DeleteArtikelById(id string) error {
	err := s.ArtikelRepo.DeleteArtikelById(id)
	return err
}
func (s ArtikelUsecaseImpl) InsertArtikel(DataArtikel model.ArtikelModel) error {
	err := tools.ValidationNotNil(DataArtikel.IdUser, DataArtikel.IdCategory, DataArtikel.Description,
		DataArtikel.ArtikelTitle, DataArtikel.Date)
	if err != nil {
		return err
	}
	err = s.ArtikelRepo.InsertArtikel(DataArtikel)
	if err != nil {
		return err
	}
	return err
}
func InitArtikelUseCase(ArtikelRepo repositories.ArtikelRepo) ArtikelUseCase {
	return &ArtikelUsecaseImpl{ArtikelRepo}
}
