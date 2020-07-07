package master

import (
	"artikel/master/controller"
	"artikel/master/repositories"
	"artikel/master/usecase"
	"database/sql"

	"github.com/gorilla/mux"
)

func InitAll(R *mux.Router, DB *sql.DB) {
	ArtikelRepo := repositories.InitArtikelRepoImpl(DB)
	ArtikelUseCase := usecase.InitArtikelUseCase(ArtikelRepo)
	controller.ArtikelController(R, ArtikelUseCase)
	UserRepo := repositories.InitUserRepoImpl(DB)
	UserUseCase := usecase.InitUserUseCase(UserRepo)
	controller.UserController(R, UserUseCase)
}
