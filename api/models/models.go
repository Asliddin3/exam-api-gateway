package models

type Error struct {
	Error string `json:"error"`
}
type CofirmEmail struct {
	UserNameOrEmail string
	Password        string
}
type AdminRequest struct {
	UserName string
	PassWord string
}
type AdminResponse struct {
	UserName    string
	AccessToken string
}

type Register struct {
	FirstName   string     `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName"`
	LastName    string     `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName"`
	Bio         string     `protobuf:"bytes,3,opt,name=Bio,proto3" json:"Bio"`
	Adderesses  []*Address `protobuf:"bytes,4,rep,name=Adderesses,proto3" json:"Adderesses"`
	Email       string     `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email"`
	PhoneNumber string     `protobuf:"bytes,6,opt,name=PhoneNumber,proto3" json:"PhoneNumber"`
	UserName    string     `protobuf:"bytes,10,opt,name=UserName,proto3" json:"UserName"`
	PassWord    string     `protobuf:"bytes,11,opt,name=PassWord,proto3" json:"PassWord"`
}

type Address struct {
	District string `protobuf:"bytes,1,opt,name=District,proto3" json:"District"`
	Street   string `protobuf:"bytes,2,opt,name=Street,proto3" json:"Street"`
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
	Id           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
