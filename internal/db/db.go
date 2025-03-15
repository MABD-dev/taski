package db

type Db interface {
	List()
	Add(name string, description string)
	Delete(number int) error
}

var (
	dbInstance Db
)

func GetDb() Db {
	if dbInstance == nil {
		dbInstance = NewPersistentDb()
	}
	return dbInstance
}
