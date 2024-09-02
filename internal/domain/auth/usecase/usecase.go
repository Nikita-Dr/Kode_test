package usecase

import (
	"Kode_test/internal/domain/auth/entity"
	"Kode_test/internal/domain/auth/model"
	"fmt"
)

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserIdByEmail(email, password string) (int, error)
}

type JWT interface {
	GenerateToken(userID int) (string, error)
	//ParseToken(tokenStr string) (int, error)
}

type AuthUseCase struct {
	userRepo UserRepository
	jwt      JWT
}

func NewAuthUseCase(userRepo UserRepository, jwt JWT) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo, jwt: jwt}
}

func (u *AuthUseCase) SignUP(userDTO model.SignUpRequestDTO) error {
	user := entity.UserFromDTO(userDTO.Email, userDTO.Password)

	if err := u.userRepo.CreateUser(user); err != nil {
		return fmt.Errorf("AuthUseCase - SignUP: %w", err)
	}

	return nil
}

func (u *AuthUseCase) Login(email, password string) (string, error) {
	id, err := u.userRepo.GetUserIdByEmail(email, password)
	if err != nil {
		return "", fmt.Errorf("AuthUseCase - GetUserUserIdByEmail: %w", err)
	}

	var token string
	if token, err = u.jwt.GenerateToken(id); err != nil {
		return "", fmt.Errorf("AuthUseCase - GenerateToken: %w", err)
	}
	return token, nil
}
