package api

import (
	"errors"
	"net/http"

	"github.com/Degreezee/labwork10/pkg/vars"

	"github.com/labstack/echo/v4"
)

// GetQuery возвращает случайное приветствие пользователю
func (srv *Server) GetCount(e echo.Context) error {
	msg, err := srv.uc.FetchCountMessage()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

// PostQuery Помещает новый вариант приветствия в БД
func (srv *Server) PostCount(e echo.Context) error {
	input := struct {
		Msg *int `json:"count"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusBadRequest, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "the 'count' field is missing")
	}

	err = srv.uc.UpdateCountMessage(*input.Msg)
	if err != nil {
		if errors.Is(err, vars.ErrorAlreadyExists) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusCreated, "Counter updated")
}
