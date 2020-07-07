package model

type UserModel struct {
	IdUser   string `json:"id_user"`
	UserName string `json:"username"`
	NoTelp   string `json;"notelp"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}
