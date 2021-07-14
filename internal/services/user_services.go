package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"tung.gallery/internal/dt/dto"
	"tung.gallery/internal/dt/entity"
	"tung.gallery/internal/middleware"
	"tung.gallery/internal/repo"
	"tung.gallery/pkg/models"
	"tung.gallery/pkg/utils"
)

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"
)

type UserServiceInterface interface {
	CreateUser(dto.UserCreateRequest) (dto.UserCreateResponse, error)
	UpdateUser(entity.Users, dto.UserUpdateRequest) (dto.UserUpdateResponse, error)
	Login(dto.UserLoginRequest) (dto.UserLoginResponse, error)
	DeleteUser(*entity.Users) (dto.UserDeleteResponse, error)
	FindUserById(uint) (*entity.Users, error)
	FindUserByEmail(dto.UserLoginRequest) (*entity.Users, error)
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
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, err.Error())
		return dto.UserCreateResponse{
				Username:     req.Username,
				Email:        req.Email,
				Password:     req.Password,
				BaseResponse: baseResponse},
			models.ErrInternalServerError
	}

	if user.Email != "" {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrEmailHasExist.Error())
		return dto.UserCreateResponse{
				Username:     req.Username,
				Email:        req.Email,
				Password:     req.Password,
				BaseResponse: baseResponse},
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
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrCreateUserFail.Error())
		return dto.UserCreateResponse{
			Username:     req.Username,
			Email:        req.Email,
			Password:     req.Password,
			BaseResponse: baseResponse}, err
	}

	baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrCreateUserFail.Error())
	res := dto.UserCreateResponse{
		Username:     req.Username,
		Email:        req.Email,
		Password:     req.Password,
		BaseResponse: baseResponse}
	return res, nil
}

func (s *userService) UpdateUser(oldUser entity.Users, req dto.UserUpdateRequest) (dto.UserUpdateResponse, error) {
	if req.Password != "" {
		err := bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(req.Password))
		if err != bcrypt.ErrMismatchedHashAndPassword {
			baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrInvalidPassword.Error())
			return dto.UserUpdateResponse{
				BaseResponse: baseResponse,
			}, err
		}
	}

	if req.Email != "" {
		user, err := s.Repo.ByEmail(req.Email)
		if err != nil && err != models.ErrNotFound {
			baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrInternalServerError.Error())
			return dto.UserUpdateResponse{BaseResponse: baseResponse}, err
		}
		if user.Email != "" {
			baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrEmailHasExist.Error())
			return dto.UserUpdateResponse{BaseResponse: baseResponse}, err
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
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrInternalServerError.Error())
		return dto.UserUpdateResponse{BaseResponse: baseResponse}, err
	}

	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "update user success")
	res := dto.UserUpdateResponse{
		BaseResponse: baseResponse,
	}

	return res, nil
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

func (s *userService) DeleteUser(user *entity.Users) (dto.UserDeleteResponse, error) {
	err := s.Repo.Delete(user.ID)
	if err != nil {
		baseResponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrInternalServerError.Error())
		return dto.UserDeleteResponse{BaseResponse: baseResponse}, err
	}

	baseResponse := utils.BaseResponse(false, AlertLvlSuccess, "delete account success")
	res := dto.UserDeleteResponse{BaseResponse: baseResponse}
	return res, nil
}

func (s *userService) Login(req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, err := s.FindUserByEmail(req)
	if errors.Is(err, models.ErrInvalidPassword) {
		baseReponse := utils.BaseResponse(false, AlertLvlInfo, models.ErrInvalidPassword.Error())
		return dto.UserLoginResponse{BaseResponse: baseReponse}, err
	}

	token := middleware.JWTAuthService().GenerateToken(user.Email, true)

	res := dto.UserLoginResponse{Token: token}
	return res, nil
}
