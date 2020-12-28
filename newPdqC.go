package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type pdqCForm struct {
	Cid              int
	AssessmentNumber string
	AssessmentDate   string
	PdqC1            int
	PdqC2            int
	PdqC3            int
	PdqC4            int
	PdqC5            int
	PdqC6            int
	PdqC7            int
	PdqC8            int
	PdqC9            int
	PdqC10           int
	PdqC11           int
	PdqC12           int
	PdqC13           int
	PdqC14           int
	PdqC15           int
	PdqC16           int
	PdqC17           int
	PdqC18           int
	PdqC19           int
	PdqC20           int
	PdqC21           int
	PdqC22           int
	PdqC23           int
	PdqC24           int
	PdqC25           int
	PdqC26           int
	PdqC27           int
	PdqC28           int
	PdqC29           int
}

func newPdqC(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		pdqc := pdqCForm{}
		if json.Unmarshal(body, &pdqc) != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
		}

		sqlStatement := `INSERT INTO pdq_carer(cid, assessment_number, assessment_date, pdq_c_1, pdq_c_2, pdq_c_3, pdq_c_4, pdq_c_5, pdq_c_6, pdq_c_7, pdq_c_8, pdq_c_9, pdq_c_10, pdq_c_11, pdq_c_12, pdq_c_13, pdq_c_14, pdq_c_15, pdq_c_16, pdq_c_17, pdq_c_18, pdq_c_19, pdq_c_20, pdq_c_21, pdq_c_22, pdq_c_23, pdq_c_24, pdq_c_25, pdq_c_26, pdq_c_27, pdq_c_28, pdq_c_29)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32);`

		_, err := db.Exec(sqlStatement, pdqc.Cid, pdqc.AssessmentNumber, pdqc.AssessmentDate, pdqc.PdqC1, pdqc.PdqC2, pdqc.PdqC3, pdqc.PdqC4, pdqc.PdqC5, pdqc.PdqC6, pdqc.PdqC7, pdqc.PdqC8, pdqc.PdqC9, pdqc.PdqC10, pdqc.PdqC11, pdqc.PdqC12, pdqc.PdqC13, pdqc.PdqC14, pdqc.PdqC15, pdqc.PdqC16, pdqc.PdqC17, pdqc.PdqC18, pdqc.PdqC19, pdqc.PdqC20, pdqc.PdqC21, pdqc.PdqC22, pdqc.PdqC23, pdqc.PdqC24, pdqc.PdqC25, pdqc.PdqC26, pdqc.PdqC27, pdqc.PdqC28, pdqc.PdqC29)
		if err != nil {
			redisLogger(fmt.Sprintf("newPdqC() failed to write to database -- %s", err.Error()))
		}

	}
}
