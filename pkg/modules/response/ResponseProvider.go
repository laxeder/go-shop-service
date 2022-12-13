package response

import "github.com/gofiber/fiber/v2"

func Ctx(ctx *fiber.Ctx) *Response {
	return &Response{ctx: ctx}
}

func Success(status int, data ...any) *Result {
	if len(data) > 0 {
		return &Result{Status: status, Data: data[0]}
	}
	return &Result{Status: status, Data: nil}
}

func Error(status int, code, message string) *Result {
	return &Result{Status: status, Code: code, Message: message}
}

func ErrorDefault(code string) *Result {
	return Error(500, code, "Serviço indisponível no momento. Tente novamente mais tarde.")
}
