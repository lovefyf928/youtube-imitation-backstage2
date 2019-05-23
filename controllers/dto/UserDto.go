package dto

type UserData struct {
	Uid         interface{}
	Token		interface{}
	UserName    interface{}
	Email       interface{}
	PhoneNumber interface{}
	Password    interface{}
	Sex         interface{}
	Birthday    interface{}
	Code        interface{}
}




func Ud(uid interface{}, userName interface{}, email interface{}, phoneNumber interface{}, password interface{}, sex interface{}, birthday interface{}, code interface{}) *UserData {
	return &UserData{Uid: uid, UserName: userName, Email: email, PhoneNumber: phoneNumber, Password: password, Sex: sex, Birthday: birthday, Code: code}
}

func LoginDto(token interface{}) *UserData {
	return &UserData{Token: token}
}

func TokenSelectUsernameAndEmailDto(userName interface{}, email interface{}) *UserData {
	return &UserData{UserName: userName, Email: email}
}


