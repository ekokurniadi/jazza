package input

type AuthLoginInput struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type AuthRegisterInput struct {
	UserName          string  `json:"user_name"`
	Password          string  `json:"password"`
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
	KtpID             string  `json:"ktp_id"`
	AccountBank       string  `json:"account_bank"`
	AccountBankNumber string  `json:"account_bank_number"`
	AccountBankName   string  `json:"account_bank_name"`
	Rating            float64 `json:"rating"`
	Skill             string  `json:"skill"`
	IsActive          int     `json:"is_active"`
	Token             string  `json:"token"`
	TokenFCM          string  `json:"token_fcm"`
}
