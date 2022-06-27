package service

import (
	"errors"
	"jazza/entity"
	"jazza/input"
	"jazza/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthUserService interface {
	FindUserByPhoneNumber(loginData input.AuthLoginInput) (entity.Auth, error)
}

type authUserService struct {
	repository repository.AuthRepository
}

func NewAuthUserService(repository repository.AuthRepository) *authUserService {
	return &authUserService{repository}
}

func (s *authUserService) FindUserByPhoneNumber(loginData input.AuthLoginInput) (entity.Auth, error) {
	password := loginData.Password
	user, err := s.repository.FindUserByPhoneNumber(loginData)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan, mohon periksa kembali nomor telepon atau password anda")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("terjadi kesalahan, password dan nomor telepon anda tidak cocok")
	}

	return user, nil
}
