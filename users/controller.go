package users

import "request-handler-unit-test-example/users/dto"

type Controller interface {
	GetUsers() (dto.GetUsersResponse, error)
	CreateUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error)
}
