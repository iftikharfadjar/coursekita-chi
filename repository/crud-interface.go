package repository

type CrudInterface interface {
	Save(c interface{}) (error)
	FindAll(c string) (interface{}, error)
	DeleteByID(c []string) error
	UpdateByID(interface{}) error
	FindByID(c string) (interface{}, error)
}
