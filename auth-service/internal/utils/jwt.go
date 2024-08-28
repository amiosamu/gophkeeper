package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/amiosamu/gophkeeper/auth-service/internal/models"
)

type JwtWraper struct {
	SecretKey string
	Issuer    string
}

type jwtClaims struct {
	jwt.StandardClaims
	IdUser int64
}

func (j *JwtWraper) GenerateToken(user *models.User) (string, error) {
	claims := &jwtClaims{
		IdUser: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * 10).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JwtWraper) ValidateToken(signedToken string) (*jwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("could not parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("jwt token is expired")
	}

	return claims, nil
}
