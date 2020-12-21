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
	Nms1             bool
	Nms2             bool
	Nms3             bool
	Nms4             bool
	Nms5             bool
	Nms6             bool
	Nms7             bool
	Nms8             bool
	Nms9             bool
	Nms10            bool
	Nms11            bool
	Nms12            bool
	Nms13            bool
	Nms14            bool
	Nms15            bool
	Nms16            bool
	Nms17            bool
	Nms18            bool
	Nms19            bool
	Nms20            bool
	Nms21            bool
	Nms22            bool
	Nms23            bool
	Nms24            bool
	Nms25            bool
	Nms26            bool
	Nms27            bool
	Nms28            bool
	Nms29            bool
	Nms30            bool
	Nms31            bool
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
