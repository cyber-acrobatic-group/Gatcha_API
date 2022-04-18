package dto

import "github.com/cyber-acrobatic-group/Gatcha_API/model"

type ShowResponse struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Count    uint64 `json:"count" binding:"required"`
}

func ToShowResponse(user model.User) ShowResponse {
	return ShowResponse{
		Name:     user.Name,
		Password: user.Password,
		Count:    user.Count,
	}
}
