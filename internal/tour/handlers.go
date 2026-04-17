package tour

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTourHandler(svc TourService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := svc.List(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			w.Write([]byte("failed to list tours"))
			return
		}

		json.NewEncoder(w).Encode(list)
	}
}

func PostTourHandler(svc TourService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload Tour

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}

		svc.Insert(r.Context(), payload)
	}
}
