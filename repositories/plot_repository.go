package repositories

import (
	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"
)

type PlotRepository struct {
	db *gorm.DB
}

func NewPlotRepository(db *gorm.DB) PlotRepository {
	return PlotRepository{db}
}

func (r *PlotRepository) GetPlotsByUser(user models.User) ([]models.Plot, error) {
	var plots []models.Plot
	if err := r.db.Where("user_id = ?", user.ID).Find(&plots).Error; err != nil {
		return nil, err
	}
	return plots, nil
}

func (r *PlotRepository) GetPlotByID(id uint) (*models.Plot, error) {
	var plot models.Plot
	if err := r.db.First(&plot, id).Error; err != nil {
		return nil, err
	}
	return &plot, nil
}

func (r *PlotRepository) CreatePlot(plot *models.Plot) error {
	if err := r.db.Create(plot).Error; err != nil {
		return err
	}
	return nil
}
