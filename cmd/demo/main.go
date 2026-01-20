package main

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/interceptor/cors"
	"github.com/NARUBROWN/spine/pkg/route"
)

func main() {
	app := spine.New()

	// 생성자 등록
	app.Constructor(
		NewUserController,
	)

	// 라우트 등록, 라우터 단위 인터셉터
	app.Route(
		"GET",
		"/users/:id",
		(*UserController).GetUser,
		route.WithInterceptors(&LoggingInterceptor{}),
	)

	app.Route(
		"POST",
		"/users",
		(*UserController).CreateUser,
	)

	app.Route(
		"GET",
		"/users",
		(*UserController).GetUserQuery,
	)

	app.Interceptor(
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "OPTIONS"},
			AllowHeaders: []string{"Content-Type"},
		}),
	)

	app.Run(":8080")
}
