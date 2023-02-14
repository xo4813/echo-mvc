// user_repository.go

package repository

import (
    "github.com/go-xorm/xorm"
    "your_app/model"
)

type UserRepository struct {
    orm *xorm.Engine
}

func NewUserRepository(orm *xorm.Engine) *UserRepository {
    return &UserRepository{
        orm: orm,
    }
}

func (r *UserRepository) CreateUser(user *model.User) error {
    _, err := r.orm.Insert(user)
    return err
}

func (r UserRepository) GetUser(id int64) (model.User, error) {
    user := &model.User{ID: id}
    has, err := r.orm.Get(user)
    if !has {
        return nil, nil
    }
    return user, err
}
