package usecase

type IRepository interface {
	Add(fullUrl, tinyUrl string) error
	Get(tinyUrl string) (string, error)
	CheckIfFullUrlExists(fullUrl string) (string, error)
	CheckIfTinyUrlExists(tinyUrl string) (bool, error)
}

type TinyUrlUsecase struct {
	baseUrl         string
	repository      IRepository
	generateTinyUrl func() string
}

func InitTinyUrlUsecase(baseUrl string, repository IRepository, urlGenerator func() string) *TinyUrlUsecase {
	return &TinyUrlUsecase{baseUrl: baseUrl, repository: repository, generateTinyUrl: urlGenerator}
}

func (u *TinyUrlUsecase) Add(fullUrl string) (string, error) {
	// Check if full url already exists
	tinyUrl, err := u.repository.CheckIfFullUrlExists(fullUrl)
	if err != nil {
		return "", err
	}
	if tinyUrl != "" {
		return u.baseUrl + tinyUrl, nil
	}

	// If generated tiny url already exists -> generate again
	var newTinyUrl string
	for exists := true; exists; {
		newTinyUrl = u.generateTinyUrl()
		exists, err = u.repository.CheckIfTinyUrlExists(newTinyUrl)
		if err != nil {
			return "", err
		}
	}

	// Add generated tiny url to repository
	err = u.repository.Add(fullUrl, newTinyUrl)
	if err != nil {
		return "", err
	}

	return u.baseUrl + newTinyUrl, nil
}

func (u *TinyUrlUsecase) Get(tinyUrl string) (string, error) {
	// Trim base part
	trimmedTinyUrl := tinyUrl[len(u.baseUrl):]
	fullUrl, err := u.repository.Get(trimmedTinyUrl)
	if err != nil {
		return "", err
	}

	return fullUrl, nil
}
