package repository

type TinyUrlSqlRepository struct {
}

// TODO: implement
func (r *TinyUrlSqlRepository) Add(fullUrl, tinyUrl string) error {
	return nil
}

// TODO: implement
func (r *TinyUrlSqlRepository) Get(tinyUrl string) (string, error) {
	return "", nil
}

// TODO: implement
func (r *TinyUrlSqlRepository) CheckIfFullUrlExists(fullUrl string) (string, error) {
	return "", nil
}

// TODO: implement
func (r *TinyUrlSqlRepository) CheckIfTinyUrlExists(tinyUrl string) (bool, error) {
	return false, nil
}
