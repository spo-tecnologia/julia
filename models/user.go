package models

import (
	"time"

	"github.com/OdairPianta/julia/enums"
	"gorm.io/gorm"
)

type User struct {
	DefaultModel
	DeletedAt          gorm.DeletedAt        `gorm:"index" json:"deleted_at"`
	Name               string                `gorm:"size:255;notNull" validate:"required" json:"name"`
	Email              string                `gorm:"size:255;notNull;unique" json:"email"`
	EmailVerifiedAt    time.Time             `json:"email_verified_at"`
	Password           string                `gorm:"size:255;notNull" validate:"required" json:"password"`
	RememberToken      string                `gorm:"size:255" json:"remember_token"`
	FCMToken           string                `gorm:"size:255" json:"fcm_token"`
	Token              string                `gorm:"size:255" json:"token"`
	ResetPasswordToken string                `gorm:"size:255" json:"reset_password_token"`
	Profile            enums.UserProfileEnum `gorm:"notNull" validate:"required" json:"profile"` // 10 - Professor, 20 - Gestor, 30 - Secretário, 40 - Administrador
	Phone              string                `gorm:"size:255" json:"phone"`
}
