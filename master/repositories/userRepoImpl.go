package repositories

import (
	"artikel/master/model"
	"database/sql"
	"log"
)

type UserRepoImpl struct {
	db *sql.DB
}

func (ar UserRepoImpl) GetAllDataUser() ([]*model.UserModel, error) {
	dataUser := []*model.UserModel{}
	query := `select * from user`
	data, err := ar.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		User := model.UserModel{}
		var err = data.Scan(&User.IdUser,
			&User.UserName, &User.NoTelp,
			&User.Address, &User.Gender)
		if err != nil {
			return nil, err
		}
		dataUser = append(dataUser, &User)
	}
	return dataUser, nil
}
func (ar UserRepoImpl) GetUserById(id string) (model.UserModel, error) {
	var User model.UserModel
	query := `select * from user where id_user=?`
	err := ar.db.QueryRow(query, id).Scan(&User.IdUser,
		&User.UserName, &User.NoTelp,
		&User.Address, &User.Gender)

	return User, err
}
func (s UserRepoImpl) DeleteUserById(id string) error {
	tr, err := s.db.Begin()

	_, _ = s.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = s.GetUserById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := "delete from user where id_user=?"
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
func (ar UserRepoImpl) UpdateUserById(id string, DataUser model.UserModel) error {
	tr, err := ar.db.Begin()

	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	_, err = ar.GetUserById(id)
	if err != nil {
		tr.Rollback()
		return err
	}
	query := `update user set nama=?, notelp=?,alamat=?,
	jenis_kelamin=? where id_user=?`
	row, err := ar.db.Query(query, &DataUser.UserName,
		&DataUser.NoTelp, &DataUser.Address, &DataUser.Gender, id)
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
func (ar UserRepoImpl) InsertUser(DataUser model.UserModel) error {
	tr, err := ar.db.Begin()

	_, _ = ar.db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	query := "insert into user values (?,?,?,?,?)"
	row, err := ar.db.Query(query, &DataUser.IdUser, &DataUser.UserName,
		&DataUser.NoTelp, &DataUser.Address, &DataUser.Gender)
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
func InitUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{db}
}
