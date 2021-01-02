package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type userLogin struct {
	Username string
	Password string
}

func authenticate(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		body, _ := ioutil.ReadAll(r.Body)
		loginAttempt := userLogin{}
		if json.Unmarshal(body, &loginAttempt) != nil {
			http.Error(w, "Couldnt read post body", http.StatusBadRequest)
		}

		sqlStatement := `SELECT * FROM users WHERE username=$1;`
		result, err := db.Query(sqlStatement, loginAttempt.Username)
		if err != nil {
			redisLogger(fmt.Sprintf("authenticate() failed -- %s", err.Error()))
		}
		defer result.Close()

		actualUser := userLogin{}
		for result.Next() {
			err2 := result.Scan(&actualUser.Username, &actualUser.Password)
			if err2 != nil {
				redisLogger(fmt.Sprintf("authenticate() failed, couldn't Scan from db -- %s", err2.Error()))
			}
		}

		// check passwords
		err3 := bcrypt.CompareHashAndPassword([]byte(actualUser.Password), []byte(loginAttempt.Password))
		if err3 != nil {
			http.Error(w, "Coudln't authenticate", http.StatusBadRequest)
			redisLogger(fmt.Sprintf("authenticate couldn't bcrypt password attempt %s", err3.Error()))
		}

		// returns with status code of 200 that CUDUI is looking for
		fmt.Fprintf(w, "Authenticated")

	}
}
