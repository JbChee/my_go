package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10001, "服务内部错误")
	InvalidParams             = NewError(10002, "参数错误")
	NotFound                  = NewError(10003, "找不到")
	UnauthorizedAuthNotExist  = NewError(10004, "鉴权失败，找不到对应的appKey或appSecret")
	UnauthorizedTokenError    = NewError(10005, "鉴权失败，token错误")
	UnauthorizedTokenTimeout  = NewError(10006, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewError(10007, "鉴权失败，token生产失败")
	TooManyRequests           = NewError(10008, "请求过多")
)
type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}



func NewError(code int, msg string) * Error{
	if _, ok := codes[code]; ok{
		panic(fmt.Sprintln("错误码 %d 已经存在， 请更换一个", code))
	}
	codes[code] = msg
	return &Error{code:code, msg:msg}
}

func (e *Error) Error() string{
	return fmt.Sprintf("错误码：%v, 错误信息%v",e.code, e.msg)
}

func (e *Error) Code() int{
	return e.code
}

func (e *Error) Msg() string{
	return e.msg
}
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

func (e *Error) StatusCode() int{
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}