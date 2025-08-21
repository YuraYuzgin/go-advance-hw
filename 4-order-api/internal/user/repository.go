package user

import (
	"cart-api/pkg/db"
	"errors"
	"fmt"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (repo *UserRepository) FindByPhone(phone string) (*User, error) {
	var user User
	result := repo.database.DB.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindBySessionId(sessionId, code string) (*User, error) {
	var user User
	result := repo.database.DB.First(&user, "session_id = ?", sessionId)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	if code != user.Code {
		return nil, errors.New("wrong code")
	}

	return &user, nil
}
