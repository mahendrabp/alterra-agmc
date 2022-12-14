package mocks

import (
	"io"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/mahendrabp/alterra-agmc/day-4/utils"
)

type EchoMock struct {
	E *echo.Echo
}

func (em *EchoMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	em.E.Validator = utils.GetValidator()
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)

	return c, rec
}
