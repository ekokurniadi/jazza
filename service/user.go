package service

import (
	"jazza/entity"
	"jazza/input"
	"jazza/repository"
)

type UserService interface {
	UserServiceGetAll() ([]entity.User, error)
	UserServiceGetByID(inputID input.InputIDUser) (entity.User, error)
	UserServiceCreate(input input.UserInput) (entity.User, error)
	UserServiceUpdate(inputID input.InputIDUser, inputData input.UserInput) (entity.User, error)
	UserServiceDeleteByID(inputID input.InputIDUser) (bool, error)
}
type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}
func (s *userService) UserServiceCreate(input input.UserInput) (entity.User, error) {
	user := entity.User{}
	user.UserName = input.UserName
	user.Password = input.Password
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.PhoneNumber = input.PhoneNumber
	user.Gender = input.Gender
	user.BirthPlace = input.BirthPlace
	user.BirthDate = input.BirthDate
	user.Religion = input.Religion
	user.Province = input.Province
	user.City = input.City
	user.Address = input.Address
	user.Age = input.Age
	user.Role = input.Role
	user.ProfilePicture = input.ProfilePicture
	user.KtpID = input.KtpID
	user.AccountBank = input.AccountBank
	user.AccountBankNumber = input.AccountBankNumber
	user.AccountBankName = input.AccountBankName
	user.Rating = input.Rating
	user.Skill = input.Skill
	user.IsActive = input.IsActive
	newUser, err := s.repository.SaveUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
func (s *userService) UserServiceUpdate(inputID input.InputIDUser, inputData input.UserInput) (entity.User, error) {
	user, err := s.repository.FindByIDUser(inputID.ID)
	if err != nil {
		return user, err
	}
	user.UserName = inputData.UserName
	user.Password = inputData.Password
	user.FirstName = inputData.FirstName
	user.LastName = inputData.LastName
	user.PhoneNumber = inputData.PhoneNumber
	user.Gender = inputData.Gender
	user.BirthPlace = inputData.BirthPlace
	user.BirthDate = inputData.BirthDate
	user.Religion = inputData.Religion
	user.Province = inputData.Province
	user.City = inputData.City
	user.Address = inputData.Address
	user.Age = inputData.Age
	user.Role = inputData.Role
	user.ProfilePicture = inputData.ProfilePicture
	user.KtpID = inputData.KtpID
	user.AccountBank = inputData.AccountBank
	user.AccountBankNumber = inputData.AccountBankNumber
	user.AccountBankName = inputData.AccountBankName
	user.Rating = inputData.Rating
	user.Skill = inputData.Skill
	user.IsActive = inputData.IsActive

	updatedUser, err := s.repository.UpdateUser(user)

	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
func (s *userService) UserServiceGetByID(inputID input.InputIDUser) (entity.User, error) {
	user, err := s.repository.FindByIDUser(inputID.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (s *userService) UserServiceGetAll() ([]entity.User, error) {
	users, err := s.repository.FindAllUser()
	if err != nil {
		return users, err
	}
	return users, nil
}
func (s *userService) UserServiceDeleteByID(inputID input.InputIDUser) (bool, error) {
	_, err := s.repository.FindByIDUser(inputID.ID)
	if err != nil {
		return false, err
	}
	_, err = s.repository.DeleteByIDUser(inputID.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Generated by Micagen at 27 Juni 2022
