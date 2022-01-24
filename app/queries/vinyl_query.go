package queries

import (
	"fiber-restful-api-tutorial/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// VinylQueries struct for queries from Vinyl model
type VinylQueries struct {
	*gorm.DB
}

// GetVinyls method for getting all vinyls.
func (q *VinylQueries) GetVinyls() ([]models.Vinyl, error) {
	// Define vinyls variable
	var vinyls []models.Vinyl
	//var result map[string]interface{}

	// Query
	if err := q.Find(&vinyls).Error; err != nil {
		return vinyls, err
	}

	// Return query result
	return vinyls, nil
}

func (q *VinylQueries) GetVinyl(id uuid.UUID) (models.Vinyl, error) {
	// Define vinyl variable
	var vinyl models.Vinyl

	// Query
	if err := q.Find(&vinyl, "id = ?", id).Error; err != nil {
		return vinyl, err
	}

	// Send query result
	return vinyl, nil
}

// CreateVinyl method for creating vinyl by given Vinyl object.
func (q *VinylQueries) CreateVinyl(v *models.Vinyl) error {
	// Send query to database
	if err := q.Model(&models.Vinyl{}).Create(v).Error; err != nil {
		return err
	}

	return nil
}

// UpdateVinyl method for updating vinyl by given Vinyl object.
func (q *VinylQueries) UpdateVinyl(id uuid.UUID, b *models.Vinyl) error {
	return nil
}
