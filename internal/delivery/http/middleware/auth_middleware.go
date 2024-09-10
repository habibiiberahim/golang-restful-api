package middleware

import (
	"github.com/gofiber/fiber/v2"
	"golang-restful-api/internal/model"
	"golang-restful-api/internal/usecase"
)

func NewAuth(userUseCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT FOUND")}
		userUseCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUseCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUseCase.Log.Warnf("Failed to verify user : %v", err)
			return fiber.ErrUnauthorized
		}

		userUseCase.Log.Debugf("User : %v", auth.Name)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
