package register_user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"async-arch/auth/internal/infrastructure/contract"
	"async-arch/auth/internal/pkg/domain"
)

var (
	ErrAlreadyExists = errors.New("already exists")
)

type In struct {
	Email     string
	Role      string
	Password  string
	FirstName string
	LastName  string
}

type usecase struct {
	usersRepo domain.UserRepository
	log       contract.Log
}

func New(usersRepo domain.UserRepository, log contract.Log) *usecase {
	return &usecase{
		usersRepo: usersRepo,
		log:       log,
	}
}

func (u *usecase) Run(ctx context.Context, in In) (*domain.User, error) {
	isExists, err := u.usersRepo.Exists(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	if isExists {
		return nil, ErrAlreadyExists
	}

	hashPass, err := u.hashPass(in.Password)
	if err != nil {
		return nil, err
	}

	user, err := u.usersRepo.Upsert(ctx, domain.User{
		Email:        in.Email,
		Role:         in.Role,
		HashPassword: hashPass,
		FirstName:    in.FirstName,
		LastName:     in.LastName,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *usecase) hashPass(pass string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passHash), nil
}
