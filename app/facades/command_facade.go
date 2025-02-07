package facades

import (
	"n4a3/clean-architecture/app/domain/entity"
	"n4a3/clean-architecture/app/integrates/repository"
)

type CommandFacade interface {
}

type commandFacade struct {
	userRepository repository.UserRepository
}

func NewCommandFacade(repo repository.UserRepository) CommandFacade {
	return commandFacade{
		userRepository: repo,
	}
}

func (c *commandFacade) Insert() {

	u := entity.User{
		Name:        "",
		Username:    "",
		Email:       "",
		Point:       nil,
		UserGroupId: nil,
		UserGroup:   nil,
	}
	u.SetInserter("system")
	c.userRepository.Insert(&u)

	users := []entity.User{
		entity.User{},
		entity.User{},
	}
	c.userRepository.BulkInsert(&users)

	// Delete
	// Delete ById

	//c.userRepository.Update()
	//c.userRepository.Updates()
	//c.userRepository.UpdateWhere()
}
