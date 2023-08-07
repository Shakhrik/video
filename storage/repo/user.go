package repo

import "github.com/Shakhrik/video_task/dto"

type UserRepo interface {
	Create(req *dto.CreateUser) (string, error)
}
