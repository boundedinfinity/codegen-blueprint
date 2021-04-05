package service

import (
	"boundedinfinity/echo_blueprint/model"
	"errors"
)

var (
	NOT_IMPLEMENTED = errors.New("You forgot the implementation")
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (t *Server) HandleV1Tag(request model.TagGetRequest, response *model.TagGetResponse) error {

	// You forgot the implementation

	return NOT_IMPLEMENTED
}
