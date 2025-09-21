package repositories

type FileRepository interface {
	Save()
	GetByID()
	Delete()
}
