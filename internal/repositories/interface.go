package repositories

type Repository interface {
	Write(email string) error
}
