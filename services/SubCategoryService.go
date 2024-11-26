package services

import (
	"time"

	"github.com/pt-sinan-akbar/entities"
	"gorm.io/gorm"
)

type SubCategoryService struct {
	DB *gorm.DB
}

func NewSubCategoryService(DB *gorm.DB) SubCategoryService {
	return SubCategoryService{DB}
}

func (scs SubCategoryService) GetAll() ([]entities.SubCategory, error) {
	var obj []entities.SubCategory
	result := scs.DB.Where("deleted_at IS NULL").Preload("Category").Find(&obj)
	return obj, result.Error
}

func (scs SubCategoryService) GetByID(id int) (entities.SubCategory, error) {
	var obj entities.SubCategory
	result := scs.DB.Where("id = ? AND deleted_at IS NULL", id).Preload("Category").First(&obj)
	return obj, result.Error
}

func (scs SubCategoryService) CreateAsync(subcategory entities.SubCategory) error {
	result := scs.DB.Create(&subcategory)
	return result.Error
}

func (scs SubCategoryService) DeleteAsync(id int) error {
	tx := scs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	obj, err := scs.GetByID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	obj.UpdatedAt = time.Now()
	obj.DeletedAt = &obj.UpdatedAt

	if err := tx.Save(&obj).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (scs SubCategoryService) EditAsync(id int, obj *entities.SubCategory) error {
	tx := scs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	oldObj, err := scs.GetByID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&oldObj).Updates(map[string]interface{}{
		"name": obj.Name,
		"logo": obj.Logo,
		"category_id": obj.CategoryId,
		"updated_at": time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}