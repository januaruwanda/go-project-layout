package presenters

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	ErrInternalServerErrorCode            = 500
	ErrorBadRequestCode                   = 400
	ErrUserCodeNotFound                   = 1000
	ErrUserCodeEmptyUsername              = 1001
	ErrUserCodeEmptyPassword              = 1002
	ErrUserCodeShortPassword              = 1003
	ErrUserCreateCodeEmptyName            = 1004
	ErrUserCreateCodeNameAlreadyExist     = 1005
	ErrUserCreateCodeUsernameAlreadyExist = 1006
	ErrUserLoginCodeUsernameNotFound      = 1007
	ErrUserLoginCodePasswordNotMatch      = 1008
	ErrCCTVCodeInvalidName                = 2001
	ErrCCTVCodeEmptyLocation              = 2002
	ErrCCTVCodeEmptyAddress               = 2003
	ErrCCTVCodeEmptyCity                  = 2004
	ErrCCTVCodeEmptyProjectName           = 2005
	ErrCCTVCreateCodeAddressAlreadyExist  = 2006
	ErrCCTVNotFound                       = 2007
	ErrObjectStoreCodeNotJPEG             = 2008
	ErrObjectStoreCodeInvalidImage        = 2009
	ErrDateFormat                         = 7171
	ErrAnalyticCodeInvalidCCTVID          = 3001
	ErrAnalyticCodeEmptyAnalytic          = 3002
	ErrAnalyticCodeEmptyConfiguration     = 3003
)

var errorMessages = map[int]string{
	ErrInternalServerErrorCode:            "Internal server error",
	ErrorBadRequestCode:                   "Bad request format",
	ErrUserCodeNotFound:                   "Username not found",
	ErrUserCodeEmptyUsername:              "Username cannot be empty.",
	ErrUserCodeEmptyPassword:              "Password cannot be empty.",
	ErrUserCodeShortPassword:              "Password is too short.",
	ErrUserCreateCodeEmptyName:            "Name cannot be empty.",
	ErrUserCreateCodeNameAlreadyExist:     "Name already exists.",
	ErrUserCreateCodeUsernameAlreadyExist: "Username already exists.",
	ErrUserLoginCodeUsernameNotFound:      "Username not found.",
	ErrUserLoginCodePasswordNotMatch:      "Password does not match.",
	ErrCCTVCodeInvalidName:                "CCTV name cannot be empty.",
	ErrCCTVCodeEmptyLocation:              "CCTV location cannot be empty.",
	ErrCCTVCodeEmptyAddress:               "CCTV address cannot be empty.",
	ErrCCTVCodeEmptyCity:                  "CCTV city cannot be empty.",
	ErrCCTVCodeEmptyProjectName:           "CCTV project name cannot be empty.",
	ErrCCTVCreateCodeAddressAlreadyExist:  "CCTV address already exist.",
	ErrCCTVNotFound:                       "CCTV by that id not found.",
	ErrObjectStoreCodeNotJPEG:             "Upload type must be image.",
	ErrObjectStoreCodeInvalidImage:        "Upload image is invalid.",
	ErrDateFormat:                         "Error Date Format",
	ErrAnalyticCodeInvalidCCTVID:          "CCTV ID must be a positive integer.",
	ErrAnalyticCodeEmptyAnalytic:          "Analytic cannot be empty.",
	ErrAnalyticCodeEmptyConfiguration:     "Configuration cannot be empty.",
}

type ErrorStatus struct {
	Code    int
	Message string
}

func NewErrorStatus(code int, err error) *ErrorStatus {
	message := errorMessages[code]
	if message == "" {
		message = "Unknown error."
	}
	log.Error().Err(err).Str("error_code", strconv.Itoa(code)).Msg(message)

	return &ErrorStatus{Code: code, Message: message}
}

func (ve *ErrorStatus) Error() string {
	return fmt.Sprintf("Code: %d, Error: %s", ve.Code, ve.Message)
}

func (ve *ErrorStatus) GetHTTPCode() int {
	if ve.Code <= 500 {
		return ve.Code
	}
	return 400
}
