package dto

type CreateUserRequest struct {
	Name string `json:"name"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type UserData struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetUsersResponse struct {
	MessageResponse
	Data []UserData `json:"data"`
}

type CreateUserResponse struct {
	MessageResponse
	Data UserData `json:"data"`
}
