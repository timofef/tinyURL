package usecase

import "tinyURL/internal/pkg/tinyURL/repository"

type TinyUrlUsecase struct {
	Repository repository.IRepository
	GenerateTinyUrl  func() string
}

var baseUrl = "http://mybaseurl.com/"

func (u *TinyUrlUsecase) Add(fullUrl string) (string, error) {
	// Check if full url already exists
	tinyUrl, err := u.Repository.CheckIfFullUrlExists(fullUrl)
	if err != nil {
		return "", err
	}
	if tinyUrl != "" {
		return baseUrl + tinyUrl, nil
	}

	// If generated tiny url already exists -> generate again
	var newTinyUrl string
	for exists := true; exists {
		newTinyUrl = u.GenerateTinyUrl()
		exists, err = u.Repository.CheckIfTinyUrlExists(newTinyUrl)
		if err != nil {
			return "", err
		}
	}

	// Add generated tiny url to repository
	err = u.Repository.Add(fullUrl, newTinyUrl)
	if err != nil {
		return "", err
	}

	return baseUrl + newTinyUrl, nil
}

// TODO: implement
func (u *TinyUrlUsecase) Get(tinyUrl string) (string, error) {

	return "", nil
}
