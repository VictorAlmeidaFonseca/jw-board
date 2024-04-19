package handler

import (
	application "github.com/VictorAlmeidaFonseca/jw-board/internal/application/services"
	"github.com/VictorAlmeidaFonseca/jw-board/internal/domain/entity"
)

type PersonHandler struct {
	service application.PersonService
}

func NewPersonHandler(service application.PersonService) *PersonHandler {
	return &PersonHandler{
		service: service,
	}
}

func (h *PersonHandler) FindAll() Response {
	data, err := h.service.FindAll()
	if err != nil {
		return Response{
			body: err.Error(),
			code: 404,
		}
	}

	return Response{
		body: data,
		code: 200,
	}
}

func (h *PersonHandler) FindByID(id int64) Response {
	data, err := h.service.FindByID(id)
	if err != nil {
		return Response{
			body: err.Error(),
			code: 404,
		}
	}

	return Response{
		body: data,
		code: 200,
	}

}

func (h *PersonHandler) Update(person entity.Person) Response {
	data, err := h.service.Update(person)
	if err != nil {
		return Response{
			body: err.Error(),
			code: 404,
		}
	}

	return Response{
		body: data,
		code: 200,
	}
}

func (h *PersonHandler) Create(person entity.Person) Response {
	data, err := h.service.Create(person)
	if err != nil {
		return Response{
			body: err.Error(),
			code: 404,
		}
	}

	return Response{
		body: data,
		code: 200,
	}
}

func (h *PersonHandler) Delete(id int64) Response {
	data, err := h.service.Delete(id)
	if err != nil {
		return Response{
			body: err.Error(),
			code: 404,
		}
	}

	return Response{
		body: data,
		code: 200,
	}
}
