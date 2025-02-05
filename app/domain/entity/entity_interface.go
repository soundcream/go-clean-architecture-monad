package entity

type Entity interface {
	GetId() int
	TableName() string
	//SetUpdater(user string)
	//SetInserter(user string)
}

type SoftDeleter interface {
	//SetDelete(user string, reason string)
	//SetRestore(user string, reason string)
}

type Activator interface {
	//SetActivate(user string)
	//SetDeActivate(user string)
}
