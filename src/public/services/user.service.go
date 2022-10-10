package services

import (
	"github.com/hibiken/asynq"
	"github.com/quanndh/go-app/adapter/dtos"
	"github.com/quanndh/go-app/adapter/queues"
	"github.com/quanndh/go-app/adapter/repositories"
	"github.com/quanndh/go-app/public/resources"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	logger         *log.Logger
	userRepository repositories.IUserRepository
	JwtService     IJwtService
	QueueClient    *asynq.Client
}

func NewUserService(logger *log.Logger, userRepository repositories.IUserRepository, jwtService IJwtService, queueClient *asynq.Client) IUserService {
	return &UserService{
		logger:         logger,
		userRepository: userRepository,
		JwtService:     jwtService,
		QueueClient:    queueClient,
	}
}

func (s UserService) CreateUser(data dtos.SignupDto) (*resources.UserResource, error) {

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	if errHash != nil {
		return nil, errHash
	}

	data.Password = string(hashedPassword)

	user, err := s.userRepository.CreateUser(data)

	if err != nil {
		return nil, err
	}

	task, err := queues.NewCreatedUserTask(user.ID)

	if err != nil {
		return nil, err
	}

	_, err = s.QueueClient.Enqueue(task)

	if err != nil {
		return nil, err
	}

	return resources.NewUserResource(user), nil
}

func (s UserService) Login(data dtos.LoginDto) (*resources.LoginResource, error) {
	user, err := s.userRepository.FindByUsername(data.Username)
	if err != nil {
		return nil, err
	}

	payload := resources.NewUserResource(user)

	token, err1 := s.JwtService.Generate(payload)

	if err1 != nil {
		return nil, err1
	}

	return resources.NewLoginResource(user, token), nil
}

func (s UserService) FindById(id uint) (*resources.UserResource, error) {
	user, err := s.userRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return resources.NewUserResource(user), nil
}
