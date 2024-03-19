package repositories

import (
	"database/sql"

	"github.com/willian-lemann/poc-golang-api/internal/entities"
)

type UsersRepository interface {
	ListAll() ([]entities.User, error)
	Get() (*entities.User, error)
}

type UsersDatabaseRepository struct {
	db *sql.DB
}

func NewUsersDatabaseRepository(db *sql.DB) *UsersDatabaseRepository {
	return &UsersDatabaseRepository{db}
}

func (u UsersDatabaseRepository) ListAll() ([]entities.User, error) {
	user := entities.User{
		Name:     "Willian Leman 3",
		Email:    "willianleman@gmail.com",
		Password: "123",
		Id:       1,
	}

	result := []entities.User{user}

	return result, nil
}

func (u UsersDatabaseRepository) Get() (*entities.User, error) {
	user := entities.NewUser("Willian Lemann", "willianleman@gmail.com", "123")
	return user, nil
}
