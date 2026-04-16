package schedule

import "time"

type Schedule struct {
	TourId   int64     `json:"tour_id"`
	StartsAt time.Time `json:"starts_at"`
}

type ScheduleWithId struct {
	Id       int64
	Schedule Schedule
}
