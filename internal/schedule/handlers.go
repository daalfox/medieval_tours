package schedule

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetScheduleHandler(svc ScheduleService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		svc.List(r.Context())
	}
}

func PostScheduleHandler(svc ScheduleService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload Schedule

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}

		svc.Insert(r.Context(), payload)
	}
}
