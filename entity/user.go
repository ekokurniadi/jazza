package entity

type User struct {
	ID                int
	UserName          string
	Password          string
	FirstName         string
	LastName          string
	PhoneNumber       string
	Gender            string
	BirthPlace        string
	BirthDate         string
	Religion          string
	Province          string
	City              string
	Address           string
	Age               string
	Role              string
	ProfilePicture    string
	KtpID             string
	AccountBank       string
	AccountBankNumber string
	AccountBankName   string
	Rating            float64
	Skill             string
	IsActive          int
	TokenFCM          string
}
