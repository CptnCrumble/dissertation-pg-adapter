package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type patient struct {
	Pid       int
	Fname     string
	Sname     string
	Gender    string
	Diagnosis int
}

func newPatient(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		newPatient := patient{}
		if json.Unmarshal(body, &newPatient) != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
		}

		sqlStatement := `INSERT INTO patients(pid,fname,sname,gender,diagnosis)
		VALUES($1,'anon','anon',$2,$3);`
		_, sqlerr := db.Exec(sqlStatement, newPatient.Pid, newPatient.Gender, newPatient.Diagnosis)
		if sqlerr != nil {
			// TODO - dont panic
			panic(sqlerr)
		}
	}
}
