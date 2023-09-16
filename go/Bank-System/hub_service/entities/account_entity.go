package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountEntity struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	UserID uuid.UUID
	User   UserEntity
	Number string `gorm:"size:11;not null;default:null" json:"number"`
	Digit  string `gorm:"size:1;not null;default:null" json:"digit"`
}

func (AccountEntity) TableName() string {
	return "accounts"
}

func (entity *AccountEntity) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewUUID()
	entity.ID = uuid
	return
}
