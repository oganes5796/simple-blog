package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oganes5796/simple-blog/models"
	"github.com/oganes5796/simple-blog/pkg/repository"
)

const (
	salt      = "JNgrvo978jCNJuJhui*yf(je:ki)i"
	secretkey = "OJEFHG89jnu9JM(G)K*9ijiuniew"
)

type AuthService struct {
	repos repository.Authorization
}

type MyClaims struct {
	jwt.MapClaims
	UserId int
}

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{repos: repos}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generateHash(user.Password)
	return a.repos.CreateUser(user)
}

func (a *AuthService) GenerateJWT(username, password string) (string, error) {
	user, err := a.repos.GetUser(username, generateHash(password))
	if err != nil {
		return "", err
	}
	claims := MyClaims{
		jwt.MapClaims{
			"userId": user.Id,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		},
		user.Id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

func (a *AuthService) ParseJWT(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretkey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	return claims.UserId, nil
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
