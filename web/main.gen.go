package web

import (
	"boundedinfinity/echo_blueprint/model"
	"boundedinfinity/echo_blueprint/service"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	validator "gopkg.in/go-playground/validator.v9"
)

type Server struct {
	e *echo.Echo
	s *service.Server
}

func New(s *service.Server) *Server {
	return &Server{
		s: s,
	}
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (t *Server) Serve() error {
	t.e = echo.New()
	t.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	t.e.Validator = &customValidator{validator: validator.New()}

	t.e.POST("/v1/tag", t.v1TagPost)

	if err := t.e.Start(":8080"); err != nil {
		return err
	}

	return nil
}

func (t *Server) v1TagPost(c echo.Context) error {
	var request model.TagGetRequest
	var response model.TagGetResponse

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := t.s.HandleV1Tag(request, &response); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
