package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type updrsForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Speech           string
	Saliva           string
	Chewing          string
	Eating           string
	Dressing         string
	Hygiene          string
	Handwriting      string
	Hobbies          string
	Turning          string
}

func newUpdrs(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		updrs := updrsForm{}
		if json.Unmarshal(body, &updrs) != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
		}

		sqlStatement := `INSERT INTO updrs(pid,assessment_number,assessment_date,speech ,saliva , chewing , eating , dressing , hygiene , handwriting , hobbies , turning)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`

		_, err := db.Exec(sqlStatement, updrs.Pid, updrs.AssessmentNumber, updrs.AssessmentDate, updrs.Speech, updrs.Saliva, updrs.Chewing, updrs.Eating, updrs.Dressing, updrs.Hygiene, updrs.Handwriting, updrs.Hobbies, updrs.Turning)
		if err != nil {
			redisLogger(fmt.Sprintf("newUpdrs() failed to write to database -- %s", err.Error()))
		}
	}
}
