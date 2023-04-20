package repository

type IRepository interface {
	Add(fullUrl, tinyUrl string) error
	Get(tinyUrl string) (string, error)
	CheckIfFullUrlExists(fullUrl string) (string, error)
	CheckIfTinyUrlExists(tinyUrl string) (bool, error)
}
