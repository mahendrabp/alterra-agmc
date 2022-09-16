package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mahendrabp/alterra-agmc/day-4/mocks"
	"github.com/mahendrabp/alterra-agmc/day-4/models"
	"github.com/mahendrabp/alterra-agmc/day-4/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// try using mock
func TestLoginUsersByEmailAndPasswordInvalidPayload(t *testing.T) {

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, _ := echoMock.RequestMock(http.MethodPost, "/", nil)

	c.SetPath("/v1/login")

	// setup handler
	asserts := assert.New(t)

	// testing
	if asserts.EqualError(LoginUsersController(c), "code=400, message=Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag") {
		// t.Log(rec)
	}

}

// try using httptest
func TestLoginUsersByEmailAndPasswordSuccess(t *testing.T) {

	login := models.User{
		Email:    "test1@gmail.com",
		Password: "12345",
	}

	t.Run("success", func(t *testing.T) {
		jsonReq, err := json.Marshal(login)
		if err != nil {
			t.Log("error on => ", err)
		}
		t.Log(jsonReq)
		assert.NoError(t, err)

		e := echo.New()
		e.Validator = utils.GetValidator()

		req, err := http.NewRequest(echo.POST, "/v1/login", strings.NewReader(string(jsonReq)))
		// assert.NoError(t, err)
		if err != nil {
			t.Log("error on => ", err)
		}
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/auth/signup")

		err = LoginUsersController(c)
		t.Log(err)

		require.NoError(t, err)
		// // assert.Equal(t, http.StatusOK, rec.Code)

	})

}
