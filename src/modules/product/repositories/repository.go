package repositories

type Repositories interface {
	ProductRepositoryCommand
	ProductRepositoryQuery
}

type RepositoriesImpl struct {
	*ProductRepositoryCommandImpl
	*ProductRepositoryQueryImpl
}

func NewRepositoriesImpl() *RepositoriesImpl {
	return &RepositoriesImpl{
		ProductRepositoryCommandImpl: &ProductRepositoryCommandImpl{},
		ProductRepositoryQueryImpl:   &ProductRepositoryQueryImpl{},
	}
}
