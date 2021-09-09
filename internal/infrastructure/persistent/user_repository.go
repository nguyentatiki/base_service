// Implement user_repository for concrete databases
package persistent

import (
	"base_service/internal/domain/entities"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) GetUser(username string) (*entities.User, error) {
	user := entities.User{}
	err := repo.Get(&user, "SELECT * FROM users WHERE username=?", username)
	return &user, err
}
