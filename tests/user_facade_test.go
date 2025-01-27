package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(id int) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Test
func TestGetUserByID(t *testing.T) {
	repo := new(MockUserRepository)
	repo.On("FindByID", 1).Return(&entity.User{ID: 1, Name: "John"}, nil)

	uc := facades.NewUserFacade(repo)

	users, err := uc.ListUsers()

	assert.NoError(t, err)
	assert.Equal(t, "John", users)
	repo.AssertExpectations(t)
}
