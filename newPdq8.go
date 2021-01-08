package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type pdq8 struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Pdq1             string
	Pdq2             string
	Pdq3             string
	Pdq4             string
	Pdq5             string
	Pdq6             string
	Pdq7             string
	Pdq8             string
}

func newPdq8(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		pdq := pdq8{}
		err := json.Unmarshal(body, &pdq)
		if err != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
			redisLogger(fmt.Sprintf("failed to read new pdq8 post -- %s", err.Error()))
		}

		sqlStatement := `INSERT INTO pdq8(pid,assessment_number,assessment_date,  pdq8_1,  pdq8_2,  pdq8_3,  pdq8_4,  pdq8_5,  pdq8_6,  pdq8_7,  pdq8_8)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

		_, sqlError := db.Exec(sqlStatement, pdq.Pid, pdq.AssessmentNumber, pdq.AssessmentDate, pdq.Pdq1, pdq.Pdq2, pdq.Pdq3, pdq.Pdq4, pdq.Pdq5, pdq.Pdq6, pdq.Pdq7, pdq.Pdq8)
		if sqlError != nil {
			redisLogger(fmt.Sprintf("new pdq8 insertion failed -- %s", sqlError.Error()))
			http.Error(w, sqlError.Error(), http.StatusBadRequest)
		}
	}
}
