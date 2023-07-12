package repository

import (
	"database/sql"
	"your-app/database"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	// Initialize the database connection
	err := database.InitializeDB()
	if err != nil {
		return nil, err
	}

	// Create a new UserRepository instance with the database connection
	repo := &UserRepository{
		db: database.GetDB(),
	}
	return repo, nil
}

// Implement the repository methods...
