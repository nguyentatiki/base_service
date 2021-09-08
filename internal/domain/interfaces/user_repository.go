//Interface that will be implemented in infrastructure
package interfaces

import "base_service/internal/domain/entities"

type UserRepository interface {
	GetUser(username string) (*entities.User, error)
}
