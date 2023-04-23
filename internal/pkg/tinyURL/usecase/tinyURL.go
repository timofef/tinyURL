package usecase

type IRepository interface {
	Add(fullUrl, tinyUrl string) error
	Get(tinyUrl string) (string, error)
	CheckIfFullUrlExists(fullUrl string) (string, error)
	CheckIfTinyUrlExists(tinyUrl string) (bool, error)
}

type TinyUrlUsecase struct {
	BaseUrl         string
	Repository      IRepository
	GenerateTinyUrl func() string
}

func (u *TinyUrlUsecase) Add(fullUrl string) (string, error) {
	// Check if full url already exists
	tinyUrl, err := u.Repository.CheckIfFullUrlExists(fullUrl)
	if err != nil {
		return "", err
	}
	if tinyUrl != "" {
		return u.BaseUrl + tinyUrl, nil
	}

	// If generated tiny url already exists -> generate again
	var newTinyUrl string
	for exists := true; exists; {
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

	return u.BaseUrl + newTinyUrl, nil
}

func (u *TinyUrlUsecase) Get(tinyUrl string) (string, error) {
	// Trim base part
	trimmedTinyUrl := tinyUrl[len(u.BaseUrl):]
	fullUrl, err := u.Repository.Get(trimmedTinyUrl)
	if err != nil {
		return "", err
	}

	return fullUrl, nil
}
