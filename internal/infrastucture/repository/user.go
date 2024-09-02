package repository

import (
	"Kode_test/internal/domain/auth/entity"
	"Kode_test/pkg/storage/postgres"
	"fmt"
)

type UserRepository struct {
	db *postgres.Storage
}

func NewUserRepository(db *postgres.Storage) *UserRepository {
	return &UserRepository{db: db}
}

func (n *UserRepository) CreateUser(user entity.User) error {
	stmt, err := n.db.Prepare("INSERT INTO users(email, password) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("UserRepository - SignUP - Prepare:%v", err)
	}
	_, err = stmt.Exec(user.Email, user.Passwd)
	if err != nil {
		return fmt.Errorf("UserRepository - SignUP - Exec:%v", err)
	}
	return nil
}

func (n *UserRepository) GetUserIdByEmail(email, password string) (int, error) {
	row, err := n.db.Query("SELECT id FROM users WHERE email = $1 and password = $2", email, password)
	if err != nil {
		return 0, fmt.Errorf("UserRepository - GetUserByEmail - Query:%v", err)
	}
	defer row.Close()

	if !row.Next() {
		return 0, fmt.Errorf("UserRepository - GetUserByEmail - Query:%v", row.Err())
	}

	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("UserRepository - GetUserByEmail - Scan:%v", err)
	}

	return id, nil
}
