package login_user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"async-arch/auth/internal/infrastructure/contract"
	"async-arch/auth/internal/pkg/domain"
	"async-arch/auth/pkg/env"
)

type In struct {
	Email    string
	Password string
}

type usecase struct {
	usersRepo domain.UserRepository
	log       contract.Log
	env       env.JWT
}

func New(usersRepo domain.UserRepository, log contract.Log, env env.JWT) *usecase {
	return &usecase{
		usersRepo: usersRepo,
		log:       log,
		env:       env,
	}
}

func (u *usecase) Run(ctx context.Context, in In) (string, error) {
	user, err := u.usersRepo.GetByEmail(ctx, in.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(in.Password))
	if err != nil {
		return "", err
	}

	token, err := u.newToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *usecase) newToken(user *domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(u.env.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
