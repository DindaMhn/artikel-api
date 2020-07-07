package model

type ArtikelModel struct {
	IdUser          string `json:"id_user"`
	IdArtikel       string `json:"id_artikel"`
	UserName        string `json:"user_name,omitempty"`
	ArtikelTitle    string `json:"title"`
	IdCategory      string `json:"id_category,omitempty"`
	ArtikelCategory string `json:"category_name"`
	Description     string `json:"description"`
	Date            string `json:"date"`
}
