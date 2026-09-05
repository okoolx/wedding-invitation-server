package sqldb

import (
	"fmt"
	"time"
)

func initializeAttendanceTable() error {
	_, err := sqlDb.Exec(`
		CREATE TABLE IF NOT EXISTS attendance (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			side VARCHAR(10),
			name VARCHAR(20),
			meal VARCHAR(20),
			count INTEGER,
			timestamp INTEGER
		)
	`)
	return err
}

func CreateAttendance(side, name, meal string, count int) error {
	_, err := sqlDb.Exec(`
		INSERT INTO attendance (side, name, meal, count, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`, side, name, meal, count, time.Now().Unix())

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type Attendance struct {
	ID        int    `json:"id"`
	Side      string `json:"side"`
	Name      string `json:"name"`
	Meal      string `json:"meal"`
	Count     int    `json:"count"`
	Timestamp int64  `json:"timestamp"`
}

func GetAttendances() ([]Attendance, error) {
	rows, err := sqlDb.Query(`
		SELECT id, side, name, meal, count, timestamp
		FROM attendance
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendances []Attendance

	for rows.Next() {
		var attendance Attendance

		err := rows.Scan(
			&attendance.ID,
			&attendance.Side,
			&attendance.Name,
			&attendance.Meal,
			&attendance.Count,
			&attendance.Timestamp,
		)
		if err != nil {
			return nil, err
		}

		attendances = append(attendances, attendance)
	}

	return attendances, rows.Err()
}