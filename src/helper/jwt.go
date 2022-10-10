package helper

import (
	"os"
	"strconv"

	jwt "github.com/golang-jwt/jwt/v4"
)

type jwtClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(id int) (string, error) {
	claims := jwtClaims{
		ID:               strconv.Itoa(id),
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(os.Getenv("SECRET_KEY"))

	if err != nil {
		return "", err
	} else {
		return signedToken, nil
	}
}

func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)

	if err != nil {
		return "", nil
	}

	claims, ok := token.Claims.(*jwtClaims)

	if ok && token.Valid {
		return claims.ID, nil
	} else {
		return "", err
	}
}
