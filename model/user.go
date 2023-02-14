// user.go

package model

type User struct {
    ID        int64  `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
}

// UserRepository 인터페이스
type UserRepository interface {
    CreateUser(user *User) error
    GetUser(id int64) (*User, error)
}
