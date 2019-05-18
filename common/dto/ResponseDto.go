package dto

type ResponseDto struct {
	Success    bool
	StatusCode DtoStatusCode
	Msg        string
	Data       interface{}
}

type UserData struct {
	Uid         interface{}
	UserName    interface{}
	Email       interface{}
	PhoneNumber interface{}
	Password    interface{}
	Sex         interface{}
	Birthday    interface{}
	Code        interface{}
}

func NewResponseDto(success bool, statusCode DtoStatusCode, msg string, data interface{}) *ResponseDto {
	return &ResponseDto{Success: success, StatusCode: statusCode, Msg: msg, Data: data}
}

func NewSuccessResponseDto(data interface{}) *ResponseDto {
	return &ResponseDto{Success: true, StatusCode: SUCCESS, Msg: "", Data: data}
}

func NewSuccessResponseDtoNilMsg(msg string) *ResponseDto {
	return &ResponseDto{Success: true, StatusCode: SUCCESS, Msg: msg, Data: nil}
}

//todo 在controllers文件夹下新建 dto 文件夹，新建 UserDto，组装ResponseDto，不要直接在这里添加 User的东西，因为 ResponseDto是不应该知道 User的，职责单一原则
func Ud(uid interface{}, userName interface{}, email interface{}, phoneNumber interface{}, password interface{}, sex interface{}, birthday interface{}, code interface{}) *UserData {
	return &UserData{Uid: uid, UserName: userName, Email: email, PhoneNumber: phoneNumber, Password: password, Sex: sex, Birthday: birthday, Code: code}
}

type DtoStatusCode int

const (
	SUCCESS         DtoStatusCode = 200
	INTERNATL_ERROR DtoStatusCode = 500
	FORBBDIEN       DtoStatusCode = 401
	UNAUTHORIZED    DtoStatusCode = 403
)
