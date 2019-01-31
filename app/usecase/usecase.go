package usecase

type Usecase interface {
	GetSeats(args string) ([]string, error)
	GetCodes(target, source string) (int, int, error)
}
