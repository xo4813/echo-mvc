// controller/user.go

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/example/models"
	"github.com/example/service"
)

// UserController는 사용자 관련 HTTP 요청을 처리합니다.
type UserController struct {
	userService service.UserService
}

// NewUserController는 UserController 인스턴스를 반환합니다.
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser는 새로운 사용자를 생성합니다.
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetUser는 주어진 ID에 해당하는 사용자를 반환합니다.
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}
	user, err := c.userService.FindUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
