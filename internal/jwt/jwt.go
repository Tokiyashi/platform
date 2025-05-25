package jwt_auth

import (
	"fmt"
	"platform/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuth struct {
}

var secretKey = []byte("env_var")

func (JWTAuth *JWTAuth) GenerateToken(user *models.User) (string, error) {
  claims := jwt.MapClaims{
    "id":        user.Id,
    "firstname": user.FirstName,
    "lastname":  user.LastName,
    "email":     user.Email,
    "role":      user.Role,
    "exp":       time.Now().Add(time.Hour * 2).Unix(),
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(secretKey)
}


func (JWTAuth *JWTAuth) ValidateToken(tokenString string) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return secretKey, nil
  })

  if err != nil || !token.Valid {
    return nil, fmt.Errorf("invalid token")
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    return nil, fmt.Errorf("cannot parse claims")
  }

  return claims, nil
}