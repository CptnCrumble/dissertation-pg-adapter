package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type pdss struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Pdss1            string
	Pdss2            string
	Pdss3            string
	Pdss4            string
	Pdss5            string
	Pdss6            string
	Pdss7            string
	Pdss8            string
	Pdss9            string
	Pdss10           string
	Pdss11           string
	Pdss12           string
	Pdss13           string
	Pdss14           string
	Pdss15           string
}

func newPdss(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		pdss := pdss{}
		err := json.Unmarshal(body, &pdss)
		if err != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
			redisLogger(fmt.Sprintf("failed to read new pdss post -- %s", err.Error()))
		}

		sqlStatement := `INSERT INTO pdss(pid,assessment_number,assessment_date,  pdss1,  pdss2,  pdss3,  pdss4,  pdss5,  pdss6,  pdss7,  pdss8,  pdss9,  pdss10,  pdss11,  pdss12,  pdss13,  pdss14,  pdss15)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18);`

		_, sqlError := db.Exec(sqlStatement, pdss.Pid, pdss.AssessmentNumber, pdss.AssessmentDate, pdss.Pdss1, pdss.Pdss2, pdss.Pdss3, pdss.Pdss4, pdss.Pdss5, pdss.Pdss6, pdss.Pdss7, pdss.Pdss8, pdss.Pdss9, pdss.Pdss10, pdss.Pdss11, pdss.Pdss12, pdss.Pdss13, pdss.Pdss14, pdss.Pdss15)
		if sqlError != nil {
			redisLogger(fmt.Sprintf("new pdss insertion failed -- %s", sqlError.Error()))
			http.Error(w, sqlError.Error(), http.StatusBadRequest)
		}
	}
}
