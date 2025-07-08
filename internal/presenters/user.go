package presenters

import (
	"strings"
)

type UserCredentials struct {
	Username string
	Password string
}

func (uc *UserCredentials) Validate() *ErrorStatus {
	uc.Username = strings.TrimSpace(uc.Username)
	uc.Password = strings.TrimSpace(uc.Password)

	if uc.Username == "" {
		return NewErrorStatus(ErrUserCodeEmptyUsername, nil)
	}

	if uc.Password == "" {
		return NewErrorStatus(ErrUserCodeEmptyPassword, nil)
	}

	if len(uc.Password) < 8 {
		return NewErrorStatus(ErrUserCodeShortPassword, nil)
	}

	return nil
}

type UserResult struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	UserType    string `json:"user_type"`
	AccessToken string `json:"access_token"`
}

type UserCreate struct {
	UUID     string `json:"-"`
	Username string
	Name     string
	Password string
}

func (uc *UserCreate) Validate() *ErrorStatus {
	uc.UUID = strings.TrimSpace(uc.UUID)
	uc.Username = strings.TrimSpace(uc.Username)
	uc.Name = strings.TrimSpace(uc.Name)
	uc.Password = strings.TrimSpace(uc.Password)

	if uc.Username == "" {
		return NewErrorStatus(ErrUserCodeEmptyUsername, nil)
	}

	if uc.Name == "" {
		return NewErrorStatus(ErrUserCreateCodeEmptyName, nil)
	}

	if uc.Password == "" {
		return NewErrorStatus(ErrUserCodeEmptyPassword, nil)
	}

	if len(uc.Password) < 8 {
		return NewErrorStatus(ErrUserCodeShortPassword, nil)
	}

	return nil
}
