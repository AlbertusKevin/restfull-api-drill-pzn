package zsimple

import "errors"

type SimpleRepository struct {
	Error bool
}

// kita buat providernya, nama functionnya bebas, namun biasanya dimulai dengan New
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{SimpleRepository: repository}, nil
	}
}