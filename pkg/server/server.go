package server

import (
	"github.com/rs/zerolog/log"
	"github.com/labstack/echo/v4"
	"github.com/glennadjrussell/vault-scim/pkg/api"
)

type Server struct {
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

