package product

type Repo struct {
	GetProduct
}

func NewUsecase(repoProduct GetProduct) *Repo {
	return &Repo{
		GetProduct: repoProduct,
	}
}
