// main.go

package main

import (
	"github.com/example/controller"
	"github.com/example/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// 서비스 생성
	userService := service.NewUserService()

	// 컨트롤러 생성
	userController := controller.NewUserController(userService)

	// 라우팅
	e.POST("/users", userController.CreateUser)
	e.GET("/users", userController.GetUser)

	// 서버 시작
	e.Logger.Fatal(e.Start(":8080"))
}
