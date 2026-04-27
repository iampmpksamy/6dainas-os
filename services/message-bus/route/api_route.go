package route

import (
	"github.com/6dainas/6dainas-message-bus/codegen"
	"github.com/6dainas/6dainas-message-bus/service"
	jsoniter "github.com/json-iterator/go"
)

type APIRoute struct {
	services *service.Services
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewAPIRoute(services *service.Services) codegen.ServerInterface {
	return &APIRoute{
		services: services,
	}
}
