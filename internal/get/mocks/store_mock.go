package mocks

import (
	"github.com/JorgeSDiaz/My-Books/internal/utils"
	"github.com/stretchr/testify/mock"
)

type StoreMock struct {
	mock.Mock
}

func (m *StoreMock) GetBooks() ([]utils.Book, error) {
	args := m.Called()
	return args.Get(0).([]utils.Book), args.Error(1)
}
