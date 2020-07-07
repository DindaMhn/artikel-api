package repositories

import (
	"artikel/master/model"
	"database/sql"
	"log"
)

type ArtikelRepoImpl struct {
	db *sql.DB
}

func (ar ArtikelRepoImpl) GetAllDataArtikel() ([]*model.ArtikelModel, error) {
	dataArtikel := []*model.ArtikelModel{}
	query := `select artikel.id_artikel, user.id_user,user.nama, artikel.judul_artikel,
	category.category_name, artikel.deskripsi, artikel.tanggal_terbit from artikel
	join user on artikel.id_user = user.id_user join category on category.id_category=artikel.id_category`
	data, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Artikel := model.ArtikelModel{}
		var err = data.Scan(&Artikel.IdArtikel,
			&Artikel.IdUser, &Artikel.UserName, &Artikel.ArtikelTitle,
			&Artikel.ArtikelCategory, &Artikel.Description, &Artikel.Date)
		if err != nil {
			return nil, err
		}
		dataArtikel = append(dataArtikel, &Artikel)
	}
	return dataArtikel, nil
}
func (ar ArtikelRepoImpl) GetArtikelById(id string) (model.ArtikelModel, error) {
	var Artikel model.ArtikelModel
	query := `select artikel.id_artikel, user.id_user,user.nama, artikel.judul_artikel,
	category.category_name, artikel.deskripsi, artikel.tanggal_terbit from artikel
	join user on artikel.id_user = user.id_user join category on category.id_category=artikel.id_category where artikel.id_artikel=?`
	err := ar.db.QueryRow(query, id).Scan(&Artikel.IdArtikel,
		&Artikel.IdUser, &Artikel.UserName, &Artikel.ArtikelTitle,
		&Artikel.ArtikelCategory, &Artikel.Description, &Artikel.Date)

	return Artikel, err
}
func (s ArtikelRepoImpl) DeleteArtikelById(id string) error {
	tr, err := s.db.Begin()

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = s.GetArtikelById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := "delete from artikel where id_artikel=?"
	row, err := s.db.Query(query, id)
	if err != nil {
		tr.Rollback()
		return err
	} else {
		tr.Commit()
	}

	row.Close()
	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return err
}
func (ar ArtikelRepoImpl) UpdateArtikelById(id string, DataArtikel model.ArtikelModel) error {
	tr, err := ar.db.Begin()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = ar.GetArtikelById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := `update artikel set id_user=?, judul_artikel=?, deskripsi=?
	, tanggal_terbit=?, id_category=? where id_artikel=?`
	row, err := ar.db.Query(query, &DataArtikel.IdUser, &DataArtikel.ArtikelTitle,
		&DataArtikel.Description,
		&DataArtikel.Date, &DataArtikel.IdCategory, id)
	if err != nil {
		tr.Rollback()
		return err
	} else {
		tr.Commit()
	}

	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return err
}
func (ar ArtikelRepoImpl) InsertArtikel(DataArtikel model.ArtikelModel) error {
	tr, err := ar.db.Begin()

	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	query := "insert into artikel values (?,?,?,?,?,?)"
	row, err := ar.db.Query(query, &DataArtikel.IdArtikel, &DataArtikel.IdUser,
		&DataArtikel.ArtikelTitle, &DataArtikel.Description, &DataArtikel.Date,
		&DataArtikel.IdCategory)
	if err != nil {
		tr.Rollback()
		log.Fatal(err)
	} else {
		tr.Commit()
	}
	row.Close()
	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	return nil
}
func InitArtikelRepoImpl(db *sql.DB) ArtikelRepo {
	return &ArtikelRepoImpl{db}
}
