package schedule

import "time"

type Schedule struct {
	TourId   int64     `json:"tour_id"`
	Price    int       `json:"price"`
	StartsAt time.Time `json:"starts_at"`
}

type ScheduleWithId struct {
	Id       int64    `json:"id"`
	Schedule Schedule `json:"schedule"`
}
