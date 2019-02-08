package middleware

import "rzd_bot/app/usecase"

type AppLayer struct {
	App usecase.Usecase
}

func InitMiddleWares(app usecase.Usecase) AppLayer {
	return AppLayer{
		App: app,
	}
}
