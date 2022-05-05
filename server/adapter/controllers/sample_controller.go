package controllers

import (
	"net/http"

	"github.com/kobutomo/react-catchup-server/server/application/usecases"
	"github.com/labstack/echo/v4"
)

type ISampleController interface {
	Get(c echo.Context) error
}

type SampleController struct {
	sampleUsecase usecases.SampleUsecase
}

func NewSampleController(sampleUsecase usecases.SampleUsecase) ISampleController {
	return SampleController{
		sampleUsecase: sampleUsecase,
	}
}

func (s SampleController) Get(c echo.Context) error {
	req := c.Request()
	xrid := req.Header.Get(echo.HeaderXRequestID)
	name, err := s.sampleUsecase.Execute(c.Request().Context(), xrid, "n.tomoro13@gmail.com")
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, name)
}
