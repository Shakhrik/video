package service

import "github.com/Shakhrik/video_task/storage"

// "github.com/Shakhrik/video_task/storage"

type Service interface {
	UserService() UserService
}

type service struct {
	userSerivce UserService
}

func NewService(stg storage.Storage) Service {
	return service{
		userSerivce: newUser(stg),
	}
}

func (s service) UserService() UserService {
	return s.userSerivce
}

// func NewServiceFacilitySet(stg storage.Storage) *ServiceFacilitySet {
// 	return &ServiceFacilitySet{
// 		UserService: newUser(stg)}
// }
