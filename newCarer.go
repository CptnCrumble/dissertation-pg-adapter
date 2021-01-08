package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type carer struct {
	Cid          int
	Fname        string
	Sname        string
	Pid          int
	Relationship string
}

func newCarer(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		newC := carer{}
		if json.Unmarshal(body, &newC) != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
		}

		// Write into db.carers
		sqlStatementC := `INSERT INTO carers(cid,fname,sname) VALUES($1,'anon','anon');`

		// Handle bad posts attempting to write in cid = 0
		if newC.Cid != 0 {
			_, sqlErrorC := db.Exec(sqlStatementC, newC.Cid)
			if sqlErrorC != nil {
				redisLogger(fmt.Sprintf("newCarer() failed to write to carer table -- %s", sqlErrorC.Error()))
			}

			// Write into db.relationships
			sqlStatementR := `INSERT INTO relationships(pid,cid,relationship) VALUES($1,$2,$3);`
			_, sqlErrorR := db.Exec(sqlStatementR, newC.Pid, newC.Cid, newC.Relationship)
			if sqlErrorR != nil {
				redisLogger(fmt.Sprintf("newCarer() failed to write to relationships table -- %s", sqlErrorR.Error()))
				http.Error(w, sqlErrorR.Error(), http.StatusBadRequest)
			}
		} else {
			redisLogger(fmt.Sprintf("newCarer() received a cid of 0, likely bad http post -- %s", string(body)))
		}

	}
}
