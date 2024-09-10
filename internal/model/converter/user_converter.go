package converter

import (
	"golang-restful-api/internal/entity"
	"golang-restful-api/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User, accessToken string) *model.UserResponse {
	return &model.UserResponse{
		AccessToken:  accessToken,
		RefreshToken: user.Token,
	}
}
