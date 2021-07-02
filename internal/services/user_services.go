package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/middleware"
	"tung.gallery/internal/repo"
	"tung.gallery/pkg/models"
)

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"
)

type UserServiceInterface interface {
	CreateUser(dto.UserCreateRequest) (dto.UserCreateResponse, error)
	UpdateUser(entity.Users, dto.UserUpdateRequest) error
	Login(dto.UserLoginRequest) (string, error)
	FindUserById(uint) (*entity.Users, error)
	FindUserByEmail(dto.UserLoginRequest) (*entity.Users, error)
	DeleteUser(entity.Users) error
}

type userService struct {
	Repo repo.UserRepositoryInterface
}

func NewUserService(r repo.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		Repo: r,
	}
}

func (s *userService) CreateUser(req dto.UserCreateRequest) (dto.UserCreateResponse, error) {
	user, err := s.Repo.ByEmail(req.Email)

	if err != nil && err != models.ErrNotFound {
		return dto.UserCreateResponse{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
				Alert:    dto.Alert{Level: AlertLvlInfo, Message: err.Error()}},
			models.ErrInternalServerError
	}

	if user.Email != "" {
		log.Println(req)
		return dto.UserCreateResponse{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
				Alert:    dto.Alert{Level: AlertLvlInfo, Message: "email has exists"}},
			models.ErrEmailHasExist
	}

	newUser := entity.Users{
		Username: req.Username,
		Email:    req.Email,
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashPassword)
	err = s.Repo.CreateUser(newUser)

	if err != nil {
		return dto.UserCreateResponse{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
			Alert:    dto.Alert{Level: AlertLvlInfo, Message: err.Error()}}, err
	}

	return dto.UserCreateResponse{Alert: dto.Alert{Level: AlertLvlSuccess, Message: "create user succes"}}, nil
}

func (s *userService) UpdateUser(oldUser entity.Users, req dto.UserUpdateRequest) error {
	if req.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(req.Password))
		if err != bcrypt.ErrMismatchedHashAndPassword {
			return err
		}
	}

	if req.Email != "" {
		user, err := s.Repo.ByEmail(req.Email)
		if err != nil && err != models.ErrNotFound {
			return err
		}
		if user.Email != "" {
			return models.ErrEmailHasExist
		}
	}

	user := entity.Users{
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
		Birthday: req.Birthday,
		ImageURL: req.ImageURL,
	}

	err := s.Repo.Update(user)
	return err
}

func (s *userService) FindUserById(id uint) (*entity.Users, error) {
	user, err := s.Repo.ByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) FindUserByEmail(req dto.UserLoginRequest) (*entity.Users, error) {
	user, err := s.Repo.ByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, models.ErrInvalidPassword
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *userService) DeleteUser(user entity.Users) error {
	err := s.Repo.Delete(user.ID)
	return err
}

func (s *userService) Login(req dto.UserLoginRequest) (string, error) {
	user, err := s.FindUserByEmail(req)
	if err != nil {
		return "", err
	}

	token := middleware.JWTAuthService().GenerateToken(user.Email, true)
	if err != nil {
		return "", err
	}

	return token, nil
}
