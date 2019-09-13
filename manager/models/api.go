package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"micro-gateway/manager/db"
)

type Api struct {
	gorm.Model
	Name string `json:"name"`
	Path string `json:path`
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

func (model ApiModel) GetAll() (interface{}, error) {
	currentDB := db.GetDB()

	var apis []Api
	if err := currentDB.Find(&apis).Error; err != nil {
		return nil, fmt.Errorf("Error in get all apis")
	}

	return apis, nil
}

func (model ApiModel) GetOne(id string) (interface{}, error) {
	currentDB := db.GetDB()

	var api Api
	if err := currentDB.Where("id = ?", id).First(&api).Error; err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return api, nil
}

func (model ApiModel) Update(id string, payload Api) (interface{}, error) {
	currentDB := db.GetDB()

	var api Api
	if err := currentDB.Model(&api).Where("id = ?", id).Updates(payload).Error; err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return api, nil
}

func (model ApiModel) Delete(id string) (error) {
	currentDB := db.GetDB()

	var api Api
	if err := currentDB.Where("id = ?", id).Delete(&api).Error; err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}