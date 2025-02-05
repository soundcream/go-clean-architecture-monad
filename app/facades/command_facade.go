package facades

type CommandFacade interface {
}

type commandFacade struct {
}

func NewCommandFacade() CommandFacade {
	return commandFacade{}
}
