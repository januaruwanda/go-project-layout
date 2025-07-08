package user_repository

import (
	"github.com/januaruwanda/go-project-layout.git/internal/domain"
	"github.com/januaruwanda/go-project-layout.git/pkg/database"
)

type UserRepository struct {
	database database.DatabaseHandlerFunc
}

func NewUserRepository(database database.DatabaseHandlerFunc) *UserRepository {
	return &UserRepository{database}
}

func (ur *UserRepository) Register(user domain.UserInsert) error {
	query := `
		INSERT INTO users (uuid, username, name, password) 
		VALUES (?, ?, ?, ?)
	`

	err := ur.database(nil,
		true,
		query,
		user.UUID,
		user.Username,
		user.Name,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) IsUserByUsernameExists(username string) (bool, error) {
	var result []domain.User

	query := `
		SELECT username, name, password
		FROM users 
		WHERE username = ?
	`

	err := ur.database(&result,
		false,
		query,
		username,
	)

	if err != nil {
		return false, err
	}

	return len(result) > 0, nil
}

func (ur *UserRepository) IsUserByNameExists(name string) (bool, error) {
	var result []domain.User

	query := `
		SELECT username, name, password
		FROM users 
		WHERE name = ?
	`

	err := ur.database(&result,
		false,
		query,
		name,
	)

	if err != nil {
		return false, err
	}

	return len(result) > 0, nil
}

func (ur *UserRepository) GetUserByUsername(result *domain.User, username string) error {
	query := `
		SELECT 
			uuid AS UUID,
			username, 
			name, 
			password
		FROM 
			users 
		WHERE 
			username = ?
	`

	err := ur.database(&result,
		false,
		query,
		username,
	)

	return err
}
