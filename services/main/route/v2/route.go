package v2

import (
	"github.com/6dainas/6dainas-os/codegen"
	"github.com/6dainas/6dainas-os/service"
)

type CasaOS struct {
	fileUploadService *service.FileUploadService
}

func NewCasaOS() codegen.ServerInterface {
	return &CasaOS{
		fileUploadService: service.NewFileUploadService(),
	}
}
