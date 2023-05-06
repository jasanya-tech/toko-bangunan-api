package repositories

type Repositories interface {
	UserRepositoryCommand
	UserRepositoryQuery
	TokenRepositoryCommand
	TokenRepositoryQuery
}

type RepositoriesImpl struct {
	*UserRepositoryCommandImpl
	*UserRepositoryQueryImpl
	*TokenRepositoryCommandImpl
	*TokenRepositoryQueryImpl
}

func NewUserRepositoryImpl() *RepositoriesImpl {
	return &RepositoriesImpl{
		UserRepositoryCommandImpl:  &UserRepositoryCommandImpl{},
		UserRepositoryQueryImpl:    &UserRepositoryQueryImpl{},
		TokenRepositoryCommandImpl: &TokenRepositoryCommandImpl{},
		TokenRepositoryQueryImpl:   &TokenRepositoryQueryImpl{},
	}
}
