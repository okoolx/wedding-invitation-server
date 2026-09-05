package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/juhonamnam/wedding-invitation-server/sqldb"
	"github.com/juhonamnam/wedding-invitation-server/types"
)

type AttendanceHandler struct {
	http.Handler
}

func (h *AttendanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 참석 여부 저장
	if r.Method == http.MethodPost {
		var attendance types.AttendanceCreate

		err := json.NewDecoder(r.Body).Decode(&attendance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"BadRequest"}`))
			return
		}

		err = sqldb.CreateAttendance(
			attendance.Side,
			attendance.Name,
			attendance.Meal,
			attendance.Count,
		)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"InternalServerError"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true}`))
		return
	}

	// 참석자 목록 조회
	if r.Method == http.MethodGet {
		attendances, err := sqldb.GetAttendances()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"InternalServerError"}`))
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"attendances": attendances,
			"total":       len(attendances),
		})

		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"error":"Method Not Allowed"}`))
}