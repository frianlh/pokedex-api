package encrypt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// NewWithClaims is function to create token with signed method and claims
func NewWithClaims(claims jwt.MapClaims, jwtKey string) (signedString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err = token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// Parse is function to parsing token value
func Parse(auth, jwtKey string) (claims jwt.MapClaims, err error) {
	bearer := string([]rune(auth)[0:6])
	if bearer != "Bearer" {
		return nil, errors.New("invalid bearer token")
	}

	signedString := strings.ReplaceAll(auth, "Bearer ", "")
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims["exp"] == nil {
		return nil, errors.New("token expired")
	}
	if int64(claims["exp"].(float64)) <= time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}
