package formatter

import "jazza/entity"

type AuthFormatter struct {
	ID                int     `json:"id"`
	UserName          string  `json:"user_name"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	PhoneNumber       string  `json:"phone_number"`
	Gender            string  `json:"gender"`
	BirthPlace        string  `json:"birth_place"`
	BirthDate         string  `json:"birth_date"`
	Religion          string  `json:"religion"`
	Province          string  `json:"province"`
	City              string  `json:"city"`
	Address           string  `json:"address"`
	Age               string  `json:"age"`
	Role              string  `json:"role"`
	ProfilePicture    string  `json:"profile_picture"`
	AccountBank       string  `json:"account_bank"`
	AccountBankNumber string  `json:"account_bank_number"`
	AccountBankName   string  `json:"account_bank_name"`
	Rating            float64 `json:"rating"`
	Skill             string  `json:"skill"`
	IsActive          int     `json:"is_active"`
	Token             string  `json:"token"`
	TokenFCM          string  `json:"token_fcm"`
}

func FormatAuth(entity entity.Auth, token string) AuthFormatter {
	authFormatter := AuthFormatter{}
	authFormatter.ID = entity.ID
	authFormatter.UserName = entity.FirstName
	authFormatter.FirstName = entity.FirstName
	authFormatter.LastName = entity.LastName
	authFormatter.PhoneNumber = entity.PhoneNumber
	authFormatter.Gender = entity.Gender
	authFormatter.BirthPlace = entity.BirthPlace
	authFormatter.BirthPlace = entity.BirthPlace
	authFormatter.Religion = entity.Religion
	authFormatter.Province = entity.Province
	authFormatter.City = entity.City
	authFormatter.Address = entity.Address
	authFormatter.Age = entity.Age
	authFormatter.Role = entity.Role
	authFormatter.ProfilePicture = entity.ProfilePicture
	authFormatter.AccountBank = entity.AccountBank
	authFormatter.AccountBankNumber = entity.AccountBankNumber
	authFormatter.AccountBankName = entity.AccountBankName
	authFormatter.Rating = entity.Rating
	authFormatter.Skill = entity.Skill
	authFormatter.IsActive = entity.IsActive
	authFormatter.Token = token
	authFormatter.TokenFCM = entity.TokenFCM

	return authFormatter
}
