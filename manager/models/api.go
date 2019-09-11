package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"micro-gateway/manager/db"
)

type Api struct {
	gorm.Model
	Name string `json:"name"`
}

type ApiModel struct {}

func (model ApiModel) Validate(payload Api) error {
	if payload.Name == "" {
		return fmt.Errorf("Name: invalid")
	}
	return nil
}

func (model ApiModel) Save(payload Api) (interface{}, error) {
	currentDB := db.GetDB()

	if err := model.Validate(payload); err != nil {
		return nil, err
	}

	if result := currentDB.Create(&payload); result.Error != nil {
		return nil, fmt.Errorf("Erro ao salvar no banco")
	}

	return nil, nil
}