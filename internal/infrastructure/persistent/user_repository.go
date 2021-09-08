// Implement user_repository for concrete databases
package persistent

import (
	"base_service/internal/domain/entities"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE user (
    username text,
    email text,
    phonenumber text
);`

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository() (*UserRepository, error) {
	conn := os.Getenv("MYSQL_CONNECTION")
	db, err := sqlx.Connect("mysql", conn)
	return &UserRepository{db}, err
}

func (repo *UserRepository) GetUser(username string) (user *entities.User, err error) {
	err = repo.Select(&user, "SELECT TOP 1 FROM users WHERE username = $1", username)
	return user, err
}
