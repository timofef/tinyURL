package usecase

type IUsecase interface {
	Add(fullUrl string) (string, error)
	Get(tinyUrl string) (string, error)
}
