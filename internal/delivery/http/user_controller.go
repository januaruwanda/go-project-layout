package http

import (
	"github.com/januaruwanda/go-project-layout.git/internal/presenters"
	user_usecase "github.com/januaruwanda/go-project-layout.git/internal/usecases/user"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Usecase *user_usecase.UserUsecase
}

func NewUserController() *UserController {
	usecase := user_usecase.NewUserUsecase()
	return &UserController{usecase}
}

// Register a new user
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        UserCreate  body      presenters.UserCreate  true  "User registration payload"
// @Router       /api/user/register [post]
func (uc *UserController) Register(ctx *fiber.Ctx) error {
	var creds presenters.UserCreate
	if err := ctx.BodyParser(&creds); err != nil {
		respValidation := presenters.NewErrorStatus(presenters.ErrorBadRequestCode, err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": respValidation.Code,
				"message":    respValidation.Message,
				"success":    false,
			},
		)
	}

	errValidate := creds.Validate()
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": errValidate.Code,
				"message":    errValidate.Message,
				"success":    false,
			},
		)
	}

	errCreate := uc.Usecase.Register(creds)
	if errCreate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": errCreate.Code,
				"message":    errCreate.Message,
				"success":    false,
			},
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "user created successfully",
			"success": true,
		},
	)
}

// Login user and get token
// @Summary      Login user
// @Description  Authenticate user and retrieve JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        UserCredentials  body      presenters.UserCredentials  true  "User login payload"
// @Router       /api/user/login [post]
func (uc *UserController) Login(ctx *fiber.Ctx) error {
	var creds presenters.UserCredentials
	if err := ctx.BodyParser(&creds); err != nil {
		respValidation := presenters.NewErrorStatus(presenters.ErrorBadRequestCode, err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": respValidation.Code,
				"message":    respValidation.Message,
				"success":    false,
			},
		)
	}

	errValidate := creds.Validate()
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": errValidate.Code,
				"message":    errValidate.Message,
				"success":    false,
			},
		)
	}

	user, errCreate := uc.Usecase.Login(creds)
	if errCreate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"statusCode": errCreate.Code,
				"message":    errCreate.Message,
				"success":    false,
			},
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"message": "user created successfully",
			"success": true,
			"result":  user,
		},
	)
}
