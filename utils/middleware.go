package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/takeme-backend/config"
	"github.com/notblessy/takeme-backend/model"
	"github.com/sirupsen/logrus"
)

// JWTClaims :nodoc:
type JWTClaims struct {
	jwt.RegisteredClaims
	ID string
}

// JWTConfig returns claims config
func JWTConfig() echojwt.Config {
	c := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JWTClaims)
		},
		SigningKey: []byte(config.JwtSecret()),
	}

	return c
}

// GetSessionClaims returns jwt claims and error
func GetSessionClaims(c echo.Context) (*JWTClaims, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTClaims)

	if claims == nil {
		logrus.WithField("ctx", Dump(c)).Error(model.ErrUnauthorized)
		return nil, model.ErrUnauthorized
	}

	return claims, nil
}

// GenerateJwtToken returns token and error
func GenerateJwtToken(userID string) (string, error) {
	logger := logrus.WithFields(logrus.Fields{
		"user_id": userID,
	})

	claims := &JWTClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JwtSecret()))
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return t, nil
}
