package services

import (
	"github.com/pt-sinan-akbar/entities"
	"gorm.io/gorm"
	"time"
)


type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService(DB *gorm.DB) CategoryService {
	return CategoryService{DB}
}

func (cs CategoryService) GetAll() ([]entities.Category, error) {
	var obj []entities.Category
	result := cs.DB.Where("deleted_at IS NULL").Find(&obj)
	return obj, result.Error
}

func (cs CategoryService) GetByID(id int) (entities.Category, error) {
	var obj entities.Category
	result := cs.DB.Where("id = ? AND deleted_at IS NULL", id).First(&obj)
	return obj, result.Error
}

func (cs CategoryService) CreateAsync(category entities.Category) error {
	result := cs.DB.Create(&category)
	return result.Error
}

func (cs CategoryService) DeleteAsync(id int) error {
	tx := cs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	obj, err := cs.GetByID(id)
	if err != nil{
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

func (cs CategoryService) EditAsync(id int, obj *entities.Category) error {
	tx := cs.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	oldObj, err := cs.GetByID(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&oldObj).Updates(map[string]interface{}{
		"name": obj.Name,
		"logo": obj.Logo,
		"updated_at": time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}