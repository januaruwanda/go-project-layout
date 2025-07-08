package user_usecase

import (
	"errors"

	"github.com/januaruwanda/go-project-layout.git/internal/domain"
	"github.com/januaruwanda/go-project-layout.git/internal/presenters"
	user_repository "github.com/januaruwanda/go-project-layout.git/internal/repositories/user"
	"github.com/januaruwanda/go-project-layout.git/internal/utils"
	"github.com/januaruwanda/go-project-layout.git/pkg/database"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository *user_repository.UserRepository
}

func NewUserUsecase() *UserUsecase {
	databaseHandler := database.CreateDatabaseHandler()
	userRepository := user_repository.NewUserRepository(databaseHandler)
	return &UserUsecase{userRepository}
}

func (uu *UserUsecase) Register(user presenters.UserCreate) *presenters.ErrorStatus {
	isUserByUsernameExist, errorIsUserByUsernameExist := uu.userRepository.IsUserByUsernameExists(user.Username)
	if errorIsUserByUsernameExist != nil {
		return presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errorIsUserByUsernameExist)
	}

	if isUserByUsernameExist {
		return presenters.NewErrorStatus(presenters.ErrUserCreateCodeUsernameAlreadyExist, nil)
	}

	isUserByNameExist, errorIsUserByNameExist := uu.userRepository.IsUserByNameExists(user.Name)
	if errorIsUserByNameExist != nil {
		return presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errorIsUserByNameExist)
	}

	if isUserByNameExist {
		return presenters.NewErrorStatus(presenters.ErrUserCreateCodeNameAlreadyExist, nil)
	}

	hashedPassword, errHash := utils.HashPassword(user.Password)
	if errHash != nil {
		return presenters.NewErrorStatus(presenters.ErrUserCodeEmptyPassword, errHash)
	}

	userCreate := domain.UserInsert{
		UUID:     uuid.NewString(),
		Username: user.Username,
		Name:     user.Name,
		Password: hashedPassword,
	}

	errRegister := uu.userRepository.Register(userCreate)
	if errRegister != nil {
		return presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errRegister)
	}

	return nil
}

func (uu *UserUsecase) Login(user presenters.UserCredentials) (*presenters.UserResult, *presenters.ErrorStatus) {
	isUserByUsernameExist, errorIsUserByUsernameExist := uu.userRepository.IsUserByUsernameExists(user.Username)
	if errorIsUserByUsernameExist != nil {
		return nil, presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errorIsUserByUsernameExist)
	}

	if !isUserByUsernameExist {
		return nil, presenters.NewErrorStatus(presenters.ErrUserCodeNotFound, nil)
	}

	var userStored domain.User
	errGetUserByUsername := uu.userRepository.GetUserByUsername(&userStored, user.Username)

	if errGetUserByUsername != nil {
		return nil, presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errGetUserByUsername)
	}

	if userStored.Name == "" {
		return nil, presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errors.New("error name is empty"))
	}

	if !utils.VerifyPassword(userStored.Password, user.Password) {
		return nil, presenters.NewErrorStatus(presenters.ErrUserLoginCodePasswordNotMatch, nil)
	}

	accessToken, errAccessToken := utils.GetToken(userStored.UUID, userStored.Username, true)
	if errAccessToken != nil {
		return nil, presenters.NewErrorStatus(presenters.ErrInternalServerErrorCode, errAccessToken)
	}

	result := presenters.UserResult{
		UUID:        userStored.UUID,
		Name:        userStored.Name,
		UserType:    "admin",
		AccessToken: accessToken,
	}

	return &result, nil
}
