package usecase

import (
	"fmt"
	"github.com/williamkoller/go-ddd-api/internal/user/domain"
	"github.com/williamkoller/go-ddd-api/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sync"
	"time"
)

var bcryptSem = make(chan struct{}, 4)

type UserUseCase struct {
	repo  repository.UserRepository
	queue chan *domain.User
	once  sync.Once
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	uc := &UserUseCase{
		repo:  r,
		queue: make(chan *domain.User, 100000), // buffer de 100k
	}

	for i := 0; i < 20; i++ {
		go uc.startWorker()
	}
	return uc
}

func (uc *UserUseCase) Register(user *domain.User) error {
	select {
	case uc.queue <- user:
		log.Println("[Register] Usuário enfileirado com sucesso")
		return nil
	case <-time.After(200 * time.Millisecond):
		log.Println("[Register] Falha ao enfileirar: fila cheia ou lenta")
		return fmt.Errorf("sistema sobrecarregado, tente novamente mais tarde")
	}
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

func (uc *UserUseCase) startWorker() {
	uc.once.Do(func() {
		for i := 0; i < 50; i++ { // 🚀 50 workers paralelos para consumir a fila rapidamente
			go uc.workerLoop(i)
		}
	})
}

func (uc *UserUseCase) workerLoop(workerID int) {
	for user := range uc.queue {
		bcryptSem <- struct{}{}
		go func(u *domain.User, id int) {
			defer func() { <-bcryptSem }()

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Printf("[Worker %d] erro ao gerar hash: %v", id, err)
				return
			}

			u.Password = string(hashedPassword)
			if err := uc.repo.Create(u); err != nil {
				log.Printf("[Worker %d] erro ao salvar usuário: %v", id, err)
			} else {
				log.Printf("[Worker %d] usuário salvo: %s", id, u.Email)
			}
		}(user, workerID)
	}
}
