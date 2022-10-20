package model

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		// Name     string `json:"name" type:"varchar(40)"`
		Email    string `json:"email" type:"varchar(40)"`
		Password string `json:"password" type:"varchar(60)"`
	}

	UserRequest struct {
		Email    string `json:"email,omitempty" validate:"email"`
		Password string `json:"password" validate:"required,min=8"`
	}
)
