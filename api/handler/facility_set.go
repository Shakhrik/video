package handler

import "github.com/Shakhrik/video_task/dto"

type handlerFacilitySet struct {
	UserService UserService
}

type UserService interface {
	CreateUser(req *dto.CreateUser) (*dto.UserID, error)
}
