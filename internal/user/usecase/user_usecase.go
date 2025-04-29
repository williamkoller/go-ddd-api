package usecase

import (
	"github.com/williamkoller/go-ddd-api/internal/user/domain"
	"github.com/williamkoller/go-ddd-api/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return uc.repo.Create(user)
}

func (uc *UserUseCase) Login(email, password string) (*domain.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) List() ([]domain.User, error) {
	return uc.repo.List()
}

func (uc *UserUseCase) Update(id uint, updated *domain.User) error {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}

	user.Name = updated.Name
	user.Email = updated.Email

	if updated.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updated.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return uc.repo.Update(user)
}

func (uc *UserUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
