package get

import (
	"testing"

	"github.com/JorgeSDiaz/My-Books/internal/get/mocks"
	"github.com/JorgeSDiaz/My-Books/internal/utils"
	"github.com/stretchr/testify/assert"
)

var (
	books = []utils.Book{
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Year: 1937},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Year: 1954},
	}
)

type testCase struct {
	testName         string
	mockStore        func(*mocks.StoreMock)
	expectedResponse []utils.Book
	expectedError    error
}

var testCases = []testCase{
	{
		testName: "TestGetBooks - Success",
		mockStore: func(m *mocks.StoreMock) {
			m.On("GetBooks").Return(books, nil)
		},
		expectedResponse: books,
		expectedError:    nil,
	},
	{
		testName: "TestGetBooks - Error",
		mockStore: func(m *mocks.StoreMock) {
			m.On("GetBooks").Return([]utils.Book{}, assert.AnError)
		},
		expectedResponse: []utils.Book{},
		expectedError:    assert.AnError,
	},
}

func TestGetBooks(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			store := new(mocks.StoreMock)
			tc.mockStore(store)

			s := NewService(store)
			response, err := s.GetBooks()

			assert.Equal(t, tc.expectedResponse, response)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
