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
		// tasks := make([]models.Task, 0)
		// dbInstance = &InMemoryDb{
		// 	Tasks: &tasks,
		// }
		dbInstance = NewPersistentDb()
	}
	return dbInstance
}
