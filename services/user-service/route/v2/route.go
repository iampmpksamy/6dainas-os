package v2

import codegen "github.com/6dainas/6dainas-user-service/codegen/user_service"

type UserService struct{}

func NewUserService() codegen.ServerInterface {
	return &UserService{}
}
