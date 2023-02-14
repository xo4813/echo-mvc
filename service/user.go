// service/user.go

package service

import (
	"errors"

	"github.com/example/models"
)

// UserService는 사용자 서비스를 나타냅니다.
type UserService interface {
	CreateUser(user *models.User) error
	FindUserByID(id uint) (*models.User, error)
}

type userService struct {
	userModel models.UserModel
}

// NewUserService는 새 UserService 인스턴스를 반환합니다.
func NewUserService(userModel models.UserModel) UserService {
	return &userService{
		userModel: userModel,
	}
}

// CreateUser는 새로운 사용자를 생성합니다.
func (s *userService) CreateUser(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	return s.userModel.Create(user)
}

// FindUserByID는 주어진 ID에 해당하는 사용자를 검색합니다.
func (s userService) FindUserByID(id uint) (models.User, error) {
	return s.userModel.FindByID(id)
}
