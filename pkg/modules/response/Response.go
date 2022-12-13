package response

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	ctx    *fiber.Ctx
	Status int
	Data   any
}

func (r *Response) handleStatus(status int) int {

	if status == 201 {
		return fiber.StatusCreated
	}

	if status == 204 {
		return fiber.StatusNoContent
	}

	if status == 401 {
		return fiber.StatusUnauthorized
	}

	if status == 0 || status < 100 {
		return fiber.StatusInternalServerError
	}

	if status < 200 {
		return fiber.StatusContinue
	}

	if status < 300 {
		return fiber.StatusOK
	}

	if status < 400 {
		return fiber.StatusMultipleChoices
	}

	if status < 500 {
		return fiber.StatusBadRequest
	}

	if status >= 500 {
		return fiber.StatusInternalServerError
	}

	return fiber.StatusInternalServerError
}

func (r *Response) handleMaessage(result *Result) (res any) {

	res = ""

	if r.Status == fiber.StatusOK {

		if fmt.Sprintf("%v", reflect.ValueOf(result.Data).Kind()) == "slice" {
			if reflect.ValueOf(result.Data).Len() == 0 {
				res = []interface{}{}
				return
			}
		}

		if fmt.Sprintf("%v", reflect.ValueOf(result.Data).Kind()) == "ptr" && result.Data == "" {
			return
		}

		if result.Data != nil {
			res = result.Data
			return
		}

		return
	}

	if r.Status == fiber.StatusCreated || r.Status == fiber.StatusNoContent {
		return
	}

	if r.Status >= 400 {
		res = fiber.Map{"code": result.Code, "message": result.Message}
		return
	}

	return
}

func (r *Response) Result(response *Result) error {
	r.Status = r.handleStatus(response.Status)
	r.Data = r.handleMaessage(response)

	if r.Data == "" {
		return r.ctx.SendStatus(r.Status)
	}

	return r.ctx.Status(r.Status).JSON(r.Data)
}
