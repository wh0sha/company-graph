package model

import "time"

type Company struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	INN       string    `json:"inn"`
	OGRN      string    `json:"ogrn"`
	Address   string    `json:"address"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}
