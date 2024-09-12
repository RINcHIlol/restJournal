package repository

type Authorization interface {
}

type Repository struct {
	Authorization Authorization
}

func NewRepository(authorization Authorization) *Repository {
	//todo
}
