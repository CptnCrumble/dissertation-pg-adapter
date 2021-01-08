package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type pkg struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	DurationDays     int
	Bks              float32
	Dks              float32
	Fds              float32
	Pti              float32
	Ptt              float32
}

func newPkg(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		pkg := pkg{}
		err := json.Unmarshal(body, &pkg)
		if err != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
			redisLogger(fmt.Sprintf("failed to read new pkg post -- %s", err.Error()))
		}

		sqlStatement := `INSERT INTO pkg(pid,assessment_number,assessment_date, duration_days, bks, dks, fds, pti, ptt)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

		_, sqlError := db.Exec(sqlStatement, pkg.Pid, pkg.AssessmentNumber, pkg.AssessmentDate, pkg.DurationDays, pkg.Bks, pkg.Dks, pkg.Fds, pkg.Pti, pkg.Ptt)
		if sqlError != nil {
			redisLogger(fmt.Sprintf("new pkg insertion failed -- %s", sqlError.Error()))
			http.Error(w, sqlError.Error(), http.StatusBadRequest)
		}
	}
}
