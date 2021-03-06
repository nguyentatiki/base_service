// Implement user_repository for concrete databases
package persistent

import (
	"base_service/internal/domain/entities"
	"base_service/internal/domain/interfaces"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) GetUser(username string) (*entities.User, error) {
	user := entities.User{}
	err := repo.Get(&user, "SELECT * FROM user WHERE username=?", username)
	return &user, err
}
