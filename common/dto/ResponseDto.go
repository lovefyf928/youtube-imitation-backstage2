package dto

type ResponseDto struct {
	Success    bool
	StatusCode DtoStatusCode
	Msg        string
	Data       interface{}
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

type DtoStatusCode int

const (
	SUCCESS         DtoStatusCode = 200
	INTERNATL_ERROR DtoStatusCode = 500
	FORBBDIEN       DtoStatusCode = 401
	UNAUTHORIZED    DtoStatusCode = 403
)
