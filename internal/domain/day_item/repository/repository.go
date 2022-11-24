package repository

type Repository interface {
	Get()
	GetMany()
	Create()
}
