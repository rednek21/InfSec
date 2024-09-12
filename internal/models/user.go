package models

type Role string

const (
	Admin       Role = "admin"
	DefaultUser Role = "user"
	GuestUser   Role = "guest"
)

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
