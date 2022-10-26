package models

type Error struct {
	Error string `json:"error"`
}

type Register struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SendSms struct {
	PhoneNum string `json:"phone_num"`
}

type Verify struct {
	Code string `json:"code"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type VerifiedResponse struct {
	Id           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
