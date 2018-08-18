package core

type SQL interface {
}

type MySQL struct {
}

func (config MySQLConfig) CreateClient() SQL {

	return &MySQL{}
}
