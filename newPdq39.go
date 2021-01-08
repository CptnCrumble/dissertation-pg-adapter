package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type pdq39 struct {
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
	Pdq9             string
	Pdq10            string
	Pdq11            string
	Pdq12            string
	Pdq13            string
	Pdq14            string
	Pdq15            string
	Pdq16            string
	Pdq17            string
	Pdq18            string
	Pdq19            string
	Pdq20            string
	Pdq21            string
	Pdq22            string
	Pdq23            string
	Pdq24            string
	Pdq25            string
	Pdq26            string
	Pdq27            string
	Pdq28            string
	Pdq29            string
	Pdq30            string
	Pdq31            string
	Pdq32            string
	Pdq33            string
	Pdq34            string
	Pdq35            string
	Pdq36            string
	Pdq37            string
	Pdq38            string
	Pdq39            string
}

func newPdq39(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		pdq := pdq39{}
		err := json.Unmarshal(body, &pdq)
		if err != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
			redisLogger(fmt.Sprintf("failed to read new pdq39 post -- %s", err.Error()))
		}

		sqlStatement := `INSERT INTO pdq39(pid,assessment_number,assessment_date, pdq39_1, pdq39_2, pdq39_3, pdq39_4, pdq39_5, pdq39_6, pdq39_7, pdq39_8, pdq39_9, pdq39_10, pdq39_11, pdq39_12, pdq39_13, pdq39_14, pdq39_15, pdq39_16, pdq39_17, pdq39_18, pdq39_19, pdq39_20, pdq39_21, pdq39_22, pdq39_23, pdq39_24, pdq39_25, pdq39_26, pdq39_27, pdq39_28, pdq39_29, pdq39_30, pdq39_31, pdq39_32, pdq39_33, pdq39_34, pdq39_35, pdq39_36, pdq39_37, pdq39_38, pdq39_39)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42);`

		_, sqlError := db.Exec(sqlStatement, pdq.Pid, pdq.AssessmentNumber, pdq.AssessmentDate, pdq.Pdq1, pdq.Pdq2, pdq.Pdq3, pdq.Pdq4, pdq.Pdq5, pdq.Pdq6, pdq.Pdq7, pdq.Pdq8, pdq.Pdq9, pdq.Pdq10, pdq.Pdq11, pdq.Pdq12, pdq.Pdq13, pdq.Pdq14, pdq.Pdq15, pdq.Pdq16, pdq.Pdq17, pdq.Pdq18, pdq.Pdq19, pdq.Pdq20, pdq.Pdq21, pdq.Pdq22, pdq.Pdq23, pdq.Pdq24, pdq.Pdq25, pdq.Pdq26, pdq.Pdq27, pdq.Pdq28, pdq.Pdq29, pdq.Pdq30, pdq.Pdq31, pdq.Pdq32, pdq.Pdq33, pdq.Pdq34, pdq.Pdq35, pdq.Pdq36, pdq.Pdq37, pdq.Pdq38, pdq.Pdq39)
		if sqlError != nil {
			redisLogger(fmt.Sprintf("new pdq39 insertion failed -- %s", sqlError.Error()))
			http.Error(w, sqlError.Error(), http.StatusBadRequest)
		}
	}
}

func getPdq39Data(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		vars := mux.Vars(r)
		pid := vars["pid"]

		sqlStatement := fmt.Sprintf("SELECT * FROM pdq39 WHERE pid=%s;", pid)
		result, err := db.Query(sqlStatement)
		if err != nil {
			redisLogger(fmt.Sprintf("getPdq39Data() tried %s but failed -- %s", sqlStatement, err.Error()))
		}
		defer result.Close()

		data := make([]pdq39, 0)
		for result.Next() {
			var p pdq39
			var pk int
			err = result.Scan(&pk, &p.Pid, &p.AssessmentNumber, &p.AssessmentDate, &p.Pdq1, &p.Pdq2, &p.Pdq3, &p.Pdq4, &p.Pdq5, &p.Pdq6, &p.Pdq7, &p.Pdq8, &p.Pdq9, &p.Pdq10, &p.Pdq11, &p.Pdq12, &p.Pdq13, &p.Pdq14, &p.Pdq15, &p.Pdq16, &p.Pdq17, &p.Pdq18, &p.Pdq19, &p.Pdq20, &p.Pdq21, &p.Pdq22, &p.Pdq23, &p.Pdq24, &p.Pdq25, &p.Pdq26, &p.Pdq27, &p.Pdq28, &p.Pdq29, &p.Pdq30, &p.Pdq31, &p.Pdq32, &p.Pdq33, &p.Pdq34, &p.Pdq35, &p.Pdq36, &p.Pdq37, &p.Pdq38, &p.Pdq39)
			if err != nil {
				redisLogger(fmt.Sprintf("getPdq39Data() - Scanning results failed -- %s", err.Error()))
			}

			data = append(data, p)
		}
		output, err := json.Marshal(data)
		if err != nil {
			fmt.Fprintf(w, "json parsing error")
			redisLogger(fmt.Sprintf("couldn't parse pids to JSON -- %s", err.Error()))
		}
		fmt.Fprintf(w, string(output))
	}
}
