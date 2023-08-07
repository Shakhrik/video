package sqlite

import (
	"github.com/Shakhrik/video_task/storage/repo"
	"github.com/jmoiron/sqlx"
)

type videoRepo struct {
	db *sqlx.DB
}

func NewVideoRepo(db *sqlx.DB) repo.VideoRepo {
	return videoRepo{db: db}
}

// func (b videoRepo) Create(req *dto.VideoCreate) (res models.ResponseWithID, err error) {
// 	id := uuid.New()
// 	var busStop models.BusStop
// 	queryBusStop := `SELECT id, min(distance) as distance
// 					 FROM bus_stop
// 					 WHERE destination_id = $1
// 					 GROUP BY destination_id, id`

// 	query := `INSERT INTO bus(name, seat_count, start_time, end_time, destination_id, bus_stop_id)
// 			  VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

// 	err = b.db.Get(&busStop, queryBusStop, req.DestinationID)
// 	if err != nil {
// 		return models.ResponseWithID{}, err
// 	}

// 	err = b.db.Get(&res.ID, query, req.Name, req.SeatCount, req.StartTime, req.EndTime, req.DestinationID, busStop.ID)
// 	if err != nil {
// 		return models.ResponseWithID{}, err
// 	}

// 	return
// }

// func (b busRepo) GetAll(destinationId, limit, page int32) (models.Buses, error) {
// 	var count int64
// 	res := []models.Bus{}

// 	offset := (page - 1) * limit

// 	query := `WITH temp AS(
// 		select b.id as id, count(seat_number) as booked_seats_count, array_agg(seat_number) booked_seats
// 		from bus b
// 		left join bus_seat bs on b.id = bs.bus_id
// 		group by b.id
// 		)
// 		select b.id, b.start_time, b.end_time, b.name, b.seat_count,
// 		d.from_place || '-' || d.to_place AS destination_name,
// 		is_full, b.seat_count - t.booked_seats_count as remaining_seats, t.booked_seats
// 		from temp t
// 		join bus b on t.id = b.id
// 		JOIN destination d ON b.destination_id = d.id
// 		WHERE b.destination_id = $1
// 		LIMIT $2 OFFSET $3`
// 	result, err := b.db.Query(query, destinationId, limit, offset)
// 	if err != nil {
// 		return models.Buses{Buses: res}, err
// 	}
// 	for result.Next() {
// 		var a models.Bus
// 		err = result.Scan(
// 			&a.ID, &a.StartTime, &a.EndTime, &a.Name,
// 			&a.SeatCount, &a.DestinationName, &a.IsFull,
// 			&a.RemainingSeats, pq.Array(&a.BookedSeats),
// 		)
// 		// if err != nil {
// 		// 	return models.Buses{Buses: res}, err
// 		// }

// 		res = append(res, a)

// 	}

// 	queryCount := `SELECT count(1) FROM bus_stop WHERE destination_id = $1 LIMIT $2 OFFSET $3`
// 	err = b.db.Get(&count, queryCount, destinationId, limit, offset)
// 	if err != nil {
// 		return models.Buses{Buses: res}, err

// 	}
// 	return models.Buses{Buses: res, Count: count}, nil
// }

// func (b busRepo) ReserveBus(req models.ReserveBus) (res models.ResponseWithID, err error) {
// 	query := `INSERT INTO bus_seat (bus_id, user_id, seat_number) VALUES($1, $2, $3) RETURNING id`
// 	err = b.db.Get(&res.ID, query, req.BusID, req.UserID, req.SeatNumber)
// 	return
// }

// func (b busRepo) Delete(id int64) (res models.ResponseWithID, err error) {
// 	query := `DELETE FROM bus WHERE id = $1 RETURNING id`
// 	err = b.db.Get(&res.ID, query, id)
// 	return
// }

// func (b busRepo) GetAllBuses(limit, page int32) (models.BriefBuses, error) {
// 	var buses []models.BriefBus
// 	var count int64
// 	offset := (page - 1) * limit
// 	query := `SELECT id, name FROM bus WHERE destination_id is not null LIMIT $1 OFFSET $2`

// 	err := b.db.Select(&buses, query, limit, offset)
// 	if err != nil {
// 		return models.BriefBuses{}, err
// 	}

// 	err = b.db.Get(&count, `SELECT count(1) FROM bus WHERE destination_id is not null`)
// 	if err != nil {
// 		return models.BriefBuses{}, err
// 	}
// 	return models.BriefBuses{Buses: buses, Count: count}, nil
// }

// func (b busRepo) ChangeStatus(req models.ChangeStatus) (res models.ResponseWithID, err error) {
// 	query := `UPDATE bus SET bus_stop_id = $1 WHERE id = $2 RETURNING id`
// 	err = b.db.Get(&res.ID, query, req.BusStopID, req.BusID)
// 	return
// }

// func (b busRepo) GetBusStops(busID int64) (models.GetAllBusStopes, error) {
// 	busStops := []models.GetBusStop{}
// 	query := `SELECT
// 				bs.id as bus_stop_id,
// 				bs.name as bus_stop_name,
// 				d.from_place || '-' || d.to_place  as destination_name,
// 				bs.distance as bus_stop_distance,
// 				case b.bus_stop_id
// 					when bs.id then true
// 				else false
// 				end is_here
// 			  FROM bus b
// 			  JOIN destination d on b.destination_id = d.id
// 			  JOIN bus_stop bs on d.id = bs.destination_id
// 			  WHERE b.id = $1
// 			  ORDER BY bs.distance asc`
// 	err := b.db.Select(&busStops, query, busID)
// 	if err != nil {
// 		return models.GetAllBusStopes{}, err
// 	}

// 	return models.GetAllBusStopes{BusStops: busStops}, nil
// }
