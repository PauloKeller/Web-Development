package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	FirstName string    `gorm:"size:100;not null;default:null" json:"firstName"`
	LastName  string    `gorm:"size:100;not null;default:null" json:"lastName"`
	Username  string    `gorm:"size:100;not null;default:null" json:"userName"`
	Email     string    `gorm:"size:100;not null;default:null; unique" json:"email"`
	Password  string    `gorm:"size:100;not null;default:null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func (UserEntity) TableName() string {
	return "users"
}

func (entity *UserEntity) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewUUID()
	entity.ID = uuid
	return
}
