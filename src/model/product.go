package model

import "gorm.io/gorm"

type (
	Product struct {
		gorm.Model
		Name     string `json:"name" type:"varchar(40)"`
		AuthorID uint   `json:"author_id"`
	}

	ProductRequest struct {
		Name     string `json:"name"`
		AuthorID uint   `json:"author_id"`
	}
)
