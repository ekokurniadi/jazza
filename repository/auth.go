package repository

import (
	"jazza/entity"
	"jazza/input"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindUserByPhoneNumber(loginData input.AuthLoginInput) (entity.Auth, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (a *authRepository) FindUserByPhoneNumber(loginData input.AuthLoginInput) (entity.Auth, error) {
	var user entity.Auth
	err := a.db.Where("phone_number = ?", loginData.PhoneNumber).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
