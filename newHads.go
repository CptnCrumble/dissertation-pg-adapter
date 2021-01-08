package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type hads struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Tense            string
	Enjoy            string
	Fright           string
	Laugh            string
	Worry            string
	Cheer            string
	Ease             string
	Slow             string
	Butterfly        string
	Interest         string
	Restless         string
	Forward          string
	Panic            string
	Book             string
}

func newHads(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		hads := hads{}
		err := json.Unmarshal(body, &hads)
		if err != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
			redisLogger(fmt.Sprintf("failed to read new hads post -- %s", err.Error()))
		}

		sqlStatement := `INSERT INTO hads(pid,assessment_number,assessment_date, tense, enjoy, fright, laugh, worry, cheer, ease, slow, butterfly, interest, restless, forward, panic, book)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17);`

		_, sqlError := db.Exec(sqlStatement, hads.Pid, hads.AssessmentNumber, hads.AssessmentDate, hads.Tense, hads.Enjoy, hads.Fright, hads.Laugh, hads.Worry, hads.Cheer, hads.Ease, hads.Slow, hads.Butterfly, hads.Interest, hads.Restless, hads.Forward, hads.Panic, hads.Book)
		if sqlError != nil {
			redisLogger(fmt.Sprintf("new hads insertion failed -- %s", sqlError.Error()))
			http.Error(w, sqlError.Error(), http.StatusBadRequest)
		}
	}
}
