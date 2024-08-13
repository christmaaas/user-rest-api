package domain

type User struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Login     string `json:"login"`
}

type CreateUserDTO struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Login          string `json:"login"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UpdateUserDTO struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Login       string `json:"login"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
