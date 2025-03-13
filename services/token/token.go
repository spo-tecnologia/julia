package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/OdairPianta/julia/config"
	"github.com/OdairPianta/julia/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(userID uint) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	return err
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(userID), nil
	}
	return 0, fmt.Errorf("invalid token")
}

func LoginCheck(email string, password string) (string, error) {
	var err error
	u := models.User{}
	err = config.DB.Model(models.User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = verifyPassword(password, u.Password)
	if err != nil {
		return "", err
	}

	token, err := GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func LoginCheckByCpf(cpf string, password string) (string, error) {
	var err error
	u := models.User{}
	err = config.DB.Model(models.User{}).Where("cpf = ?", cpf).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = verifyPassword(password, u.Password)
	if err != nil {
		return "", err
	}

	token, err := GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
