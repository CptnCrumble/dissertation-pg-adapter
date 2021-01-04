package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type nmsForm struct {
	Pid              int
	AssessmentNumber string
	AssessmentDate   string
	Nms1             string
	Nms2             string
	Nms3             string
	Nms4             string
	Nms5             string
	Nms6             string
	Nms7             string
	Nms8             string
	Nms9             string
	Nms10            string
	Nms11            string
	Nms12            string
	Nms13            string
	Nms14            string
	Nms15            string
	Nms16            string
	Nms17            string
	Nms18            string
	Nms19            string
	Nms20            string
	Nms21            string
	Nms22            string
	Nms23            string
	Nms24            string
	Nms25            string
	Nms26            string
	Nms27            string
	Nms28            string
	Nms29            string
	Nms30            string
	Nms31            string
}

func newNms(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		nms := nmsForm{}
		if json.Unmarshal(body, &nms) != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
		}

		sqlStatement := `INSERT INTO nms(pid,assessment_number,assessment_date,nms1, nms2, nms3, nms4, nms5, nms6, nms7, nms8, nms9, nms10, nms11, nms12, nms13, nms14, nms15, nms16, nms17, nms18, nms19, nms20, nms21, nms22, nms23, nms24, nms25, nms26, nms27, nms28, nms29, nms30, nms31)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34);`

		_, sqlerr := db.Exec(sqlStatement, nms.Pid, nms.AssessmentNumber, nms.AssessmentDate, nms.Nms1, nms.Nms2, nms.Nms3, nms.Nms4, nms.Nms5, nms.Nms6, nms.Nms7, nms.Nms8, nms.Nms9, nms.Nms10, nms.Nms11, nms.Nms12, nms.Nms13, nms.Nms14, nms.Nms15, nms.Nms16, nms.Nms17, nms.Nms18, nms.Nms19, nms.Nms20, nms.Nms21, nms.Nms22, nms.Nms23, nms.Nms24, nms.Nms25, nms.Nms26, nms.Nms27, nms.Nms28, nms.Nms29, nms.Nms30, nms.Nms31)
		if sqlerr != nil {
			redisLogger(fmt.Sprintf("newNms() failed to write to database -- %s", sqlerr.Error()))
		}
	}
}
