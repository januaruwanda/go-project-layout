package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	EncryptionJWTKey string
}

var settings *Config

func InitCrypto() {
	settings = &Config{
		EncryptionJWTKey: viper.GetString("encryption.jwt_private_key"),
	}
}

func GetToken(userUUID, name string, isAdmin bool) (string, error) {

	claims := jwt.MapClaims{
		"uuid":  userUUID,
		"name":  name,
		"admin": isAdmin,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(settings.EncryptionJWTKey))
}

func GetJWTKey() string {
	return settings.EncryptionJWTKey
}

func VerifyPassword(storedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
