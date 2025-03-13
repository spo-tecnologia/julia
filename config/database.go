package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/OdairPianta/julia/enums"
	"github.com/OdairPianta/julia/models"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDatabase() {
	if DB != nil {
		return
	}
	err := godotenv.Load(".env")

	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		sentry.CaptureException(err)
		panic("Failed to connect to database!!! error:" + err.Error() + " dsn:" + dsn)
	}

	migrate(database)
	DB = database
	CreateOrUpdateAdminAccount()
}

func migrate(database *gorm.DB) {
	database.Migrator().AutoMigrate(&models.User{})
	database.Migrator().AutoMigrate(&models.SampleModel{})
	database.Migrator().AutoMigrate(&models.SampleItem{})
	database.Migrator().AutoMigrate(&models.File{})
	database.Migrator().AutoMigrate(&models.Job{})
	database.Migrator().AutoMigrate(&models.FailedJob{})
}

func CreateOrUpdateAdminAccount() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("julia2admin"), bcrypt.DefaultCost)
	if err != nil {
		panic("Failed to hahsed password!!! error:" + err.Error())
	}

	var user models.User
	DB.First(&user, "email = ?", "admin@admin.com")
	user.Name = "Admin"
	user.Email = "admin@admin.com"
	user.EmailVerifiedAt = time.Now()
	user.Password = string(hashedPassword)
	user.Profile = enums.UserProfileEnumAdministrator

	updateErr := DB.Save(&user).Error
	if updateErr != nil {
		panic("Failed to create admin user!!! error:" + updateErr.Error())
	}
}
