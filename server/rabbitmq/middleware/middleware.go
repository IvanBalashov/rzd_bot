package middleware

import "rzd_bot/app/usecase"

type EventLayer struct {
	App       usecase.Usecase
	LogChanel chan string
}

func InitMiddleWares(app usecase.Usecase, logChan chan string) EventLayer {
	return EventLayer{
		App:       app,
		LogChanel: logChan,
	}
}
