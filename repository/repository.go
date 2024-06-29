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
	result := tx.Create(&model)

	err = result.Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}

func (repo *Repository) Find(url string) (model model.Mapping, err error) {
	tx := repo.DB.Begin()
	result := tx.First(&model, "source = ?", url)
	err = result.Error
	return
}

func (repo *Repository) Delete(model model.Mapping) (err error) {
	tx := repo.DB.Begin()
	result := tx.Delete(&model)

	err = result.Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}

func (repo *Repository) Update(updated model.Mapping) (err error) {
	tx := repo.DB.Begin()
	result := tx.Model(&model.Mapping{}).Where("id = ?", updated.ID).Updates(updated)

	err = result.Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}
