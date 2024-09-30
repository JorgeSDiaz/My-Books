package get

import "github.com/JorgeSDiaz/My-Books/internal/utils"

type Store interface {
	GetBooks() ([]utils.Book, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) GetBooks() ([]utils.Book, error) {
	response, err := s.store.GetBooks()
	if err != nil {
		return []utils.Book{}, err
	}

	return response, nil
}
