package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PWORD")
	dbname := os.Getenv("PG_DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", greeter).Methods("GET")
	r.HandleFunc("/db_check", dbCheck(db)).Methods("GET")
	r.HandleFunc("/users", getAllUsers(db)).Methods("GET")
	r.HandleFunc("/new_user", newUser(db)).Methods("POST")

	serve(r)
	db.Close()
}

type patient struct {
	Pid       int
	Fname     string
	Sname     string
	Gender    string
	Diagnosis int
}

func serve(router *mux.Router) {
	err := http.ListenAndServe(":9057", router)
	if err != nil {
		log.Fatal("ListenAndServe failed ", err)
	}
}

func greeter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yeah yeah I'm working")
}

func dbCheck(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		fmt.Print("db ping failed")
		fmt.Print(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully connected")
	}
}

func getAllUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := "SELECT * FROM patients;"
		result, err := db.Query(sqlStatement)
		if err != nil {
			panic(err)
		}
		defer result.Close()

		users := make([]patient, 0)
		for result.Next() {
			var pid int
			var gender string
			var diagnosis int
			var fname string
			var sname string
			err = result.Scan(&pid, &fname, &sname, &gender, &diagnosis)
			if err != nil {
				panic(err)
			}

			users = append(users, patient{pid, "anon", "anon", gender, diagnosis})
		}
		output, err := json.Marshal(users)
		if err != nil {
			fmt.Fprintf(w, "json parsing error")
		}
		fmt.Fprintf(w, string(output))
	}
}

func newUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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
