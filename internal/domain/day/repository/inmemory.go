package repository

type InmemoryRepo struct {
}

func NewInmemory() Repository {
	return &InmemoryRepo{}
}
