package repositories

type Repositories interface {
	CategoryProductRepositoryCommand
	CategoryProductRepositoryQuery
}

type RepositoriesImpl struct {
	*CategoryProductRepositoryCommandImpl
	*CategoryProductRepositoryQueryImpl
}

func NewRepositoriesImpl() *RepositoriesImpl {
	return &RepositoriesImpl{
		CategoryProductRepositoryCommandImpl: &CategoryProductRepositoryCommandImpl{},
		CategoryProductRepositoryQueryImpl:   &CategoryProductRepositoryQueryImpl{},
	}
}
