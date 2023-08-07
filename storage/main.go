package storage

import (
	"github.com/Shakhrik/video_task/storage/repo"
	"github.com/Shakhrik/video_task/storage/sqlite"
	"github.com/jmoiron/sqlx"
)

// import (
// 	// "github.com/Shakhrik/inha/video_task/api/storage/postgres"
// 	// "github.com/Shakhrik/inha/video_task/api/storage/repo"
// 	"github.com/Shakhrik/inha/video_task/storage/postgres"
// 	"github.com/Shakhrik/inha/video_task/storage/repo"
// 	"github.com/jmoiron/sqlx"
// )

// type StorageI interface {
// 	UserRepo() repo.UserRepoI
// 	DestinationRepo() repo.DestinationRepoI
// 	BusStopRepo() repo.BusStopRepoI
// 	BusRepo() repo.BusRepoI
// }

type Storage interface {
	UserRepo() repo.UserRepo
}

type storage struct {
	userRepo repo.UserRepo
}

func NewStorage(db *sqlx.DB) Storage {
	return storage{
		userRepo: sqlite.NewUserRepo(db),
	}
}

func (s storage) UserRepo() repo.UserRepo {
	return s.userRepo
}

// func NewStoragePG(db *sqlx.DB) StorageI {
// 	return storagePG{
// 		db:              db,
// 		userRepo:        postgres.NewUserRepo(db),
// 		destinationRepo: postgres.NewDestinationRepo(db),
// 		busStopRepo:     postgres.NewBusStopRepo(db),
// 		busRepo:         postgres.NewBusRepo(db),
// 	}
// }

// func (s storagePG) UserRepo() repo.UserRepoI {
// 	return s.userRepo
// }

// func (s storagePG) DestinationRepo() repo.DestinationRepoI {
// 	return s.destinationRepo
// }

// func (s storagePG) BusStopRepo() repo.BusStopRepoI {
// 	return s.busStopRepo
// }

// func (s storagePG) BusRepo() repo.BusRepoI {
// 	return s.busRepo
// }
