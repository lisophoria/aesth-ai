package dto

type UpdateUserDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	ID uint `json:"id"`
	Email string `json:"email"`
}
