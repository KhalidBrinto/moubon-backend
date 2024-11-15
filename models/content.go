package models

import (
	"backend/utils"

	"gorm.io/gorm"
)

type ContentImage struct {
	ID         uint   `gorm:"primaryKey"`
	Position   string `gorm:"not null; check:position IN ('banner')"`
	Image      string `gorm:"-" json:"Image"`
	ImageBytes []byte `gorm:"column:image;type:bytea" json:"-"` // Binary data for the image
}

func (c *ContentImage) AfterFind(tx *gorm.DB) (err error) {
	c.Image = utils.EncodeImageToBase64(c.ImageBytes)

	return nil

}

func (c *ContentImage) BeforeCreate(tx *gorm.DB) (err error) {
	bt, err := utils.DecodeBase64Image(c.Image)
	if err != nil {
		return err
	}

	c.ImageBytes = bt

	return nil

}
