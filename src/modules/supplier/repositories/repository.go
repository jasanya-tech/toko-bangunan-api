package repositories

type Repositories struct {
	SupplierRepositoryCommand
	SupplierRepositoryQuery
}

type RepositoriesImpl interface {
	*SupplierRepositoryCommandImpl
	*SupplierRepositoryQueryImpl
}

func NewRepositoriesImpl() *Repositories {
	return &Repositories{
		SupplierRepositoryCommand: &SupplierRepositoryCommandImpl{},
		SupplierRepositoryQuery:   &SupplierRepositoryQueryImpl{},
	}
}
