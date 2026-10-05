package model

type Company struct {
	Name   string `json:"name"`
	INN    string `json:"inn"`
	OGRN   string `json:"ogrn"`
	Source string `json:"source"`
}
