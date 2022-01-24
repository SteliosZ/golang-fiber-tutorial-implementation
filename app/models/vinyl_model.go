package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

// Vinyl struct to describe vinyl object.
type Vinyl struct {
	gorm.Model  `json:"vinyl"`
	ID          uuid.UUID  `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	UserID      uuid.UUID  `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title       string     `db:"title" json:"title" validate:"required,lte=255"`
	Artist      string     `db:"artist" json:"artist" validate:"required,lte=255"`
	VinylStatus int        `db:"vinyl_status" json:"vinyl_status" validate:"required,len=1"`
	VinylAttrs  VinylAttrs `db:"vinyl_attrs" json:"vinyl_attrs" validate:"required,dive"`
}

// VinylAttrs struct to describe vinyl attributes.
type VinylAttrs struct {
	Cover       string `json:"cover"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}

// Value make the VinylAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b VinylAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan make the VinylAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *VinylAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}
