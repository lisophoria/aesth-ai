package repositories

import (
	"aesth-api/models"

	"gorm.io/gorm"
)

type RelevancyRepository struct {
	db *gorm.DB
}

func NewRelevancyRepository(db *gorm.DB) *RelevancyRepository {
	return &RelevancyRepository{db}
}


func (r *RelevancyRepository) Create(relevancy *models.Relevancy) error {
	return r.db.Create(relevancy).Error
}

func (r *RelevancyRepository) GetMany() ([]models.Relevancy, error) {
	var relevancys []models.Relevancy
	err := r.db.Find(&relevancys).Error
	return relevancys, err  
}

func (r *RelevancyRepository) GetByUserID(userID uint) (*models.Relevancy, error) {
	var relevancy models.Relevancy
	result := r.db.Where("user_id = ?", userID).First(&relevancy)
	if result.Error != nil {
		return nil, result.Error
	}
	return &relevancy, nil
}

func (r *RelevancyRepository) Update(relevancy *models.Relevancy) error {
	return r.db.Save(relevancy).Error
}

func (r *RelevancyRepository) Delete(id uint) error {
	return r.db.Delete(&models.Relevancy{}, id).Error
}