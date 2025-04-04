package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"n4a3/clean-architecture/app/core"
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/facades"
	"n4a3/clean-architecture/app/integrates/db"
	"testing"
)

type MockUserRepository struct {
	db.ReadOnlyRepository[entity.User]
	db.Repository[entity.User]
	mock.Mock
}

func (m *MockUserRepository) GetSpecialLogicUser(id int) *entity.User {
	args := m.Called(id)
	return args.Get(0).(*entity.User)
}

// Test
func TestGetUserByID(t *testing.T) {
	repo := new(MockRepository[entity.User])
	userRepo := new(MockUserRepository)
	userRepo.ReadOnlyRepository = repo
	repo.On("FindById", 1).Return(core.RightEither[entity.User, core.ErrContext](entity.User{
		BaseEntity: entity.NewBaseWithId(3),
		Name:       "John"}))

	facade := facades.NewUserFacade(userRepo)
	u, err := facade.GetUserById(1)
	//users, err := uc.ListUsers()

	assert.NoError(t, err)
	assert.Equal(t, "John", u.Name)
	repo.AssertExpectations(t)
}
