package repository

import (
	model "github.com/kaburov38/GoShort/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) Insert(model model.Mapping) (err error) {
	tx := repo.DB.Begin()
	tx.Create(&model)

	err = tx.Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}

func (repo *Repository) Find(url string) (model model.Mapping, err error) {
	tx := repo.DB.Begin()
	tx.First(&model, "source = ?", url)
	err = tx.Error
	return
}

func (repo *Repository) Delete(model model.Mapping) (err error) {
	tx := repo.DB.Begin()
	tx.Delete(&model)

	err = tx.Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}

func (repo *Repository) Update(old model.Mapping, new model.Mapping) (err error) {
	tx := repo.DB.Begin()
	tx.Model(&old).Updates(&new)

	err = tx.Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}
