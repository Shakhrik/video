package service

import (
	"fmt"

	"github.com/Shakhrik/video_task/dto"
	st "github.com/Shakhrik/video_task/storage"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req *dto.CreateUser) (*dto.UserID, error)
}

type User struct {
	stg st.Storage
}

func newUser(stg st.Storage) *User {
	return &User{stg: stg}
}

func (u User) CreateUser(req *dto.CreateUser) (*dto.UserID, error) {
	id := uuid.New()
	req.ID = id.String()
	uID, err := u.stg.UserRepo().Create(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return &dto.UserID{ID: uID}, nil
}
