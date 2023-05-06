package repositories

type Repositories interface {
	PurchaseProductRepositoryCommand
	PurchaseProductRepositoryQuery
}

type RepositoriesImpl struct {
	*PurchaseProductRepositoryCommandImpl
	*PurchaseProductRepositoryQueryImpl
}

func NewRepositoriesImpl() *RepositoriesImpl {
	return &RepositoriesImpl{
		PurchaseProductRepositoryCommandImpl: &PurchaseProductRepositoryCommandImpl{},
		PurchaseProductRepositoryQueryImpl:   &PurchaseProductRepositoryQueryImpl{},
	}
}
