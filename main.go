package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	// r.HandleFunc("/new_user", newUser(db)).Methods("POST")

	serve(r)
	db.Close()
}

type client struct {
	Id        int
	Role      string
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
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Successfully connected")
	}
}

func getAllUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sqlStatement := "SELECT * FROM public.users;"
		result, err := db.Query(sqlStatement)
		if err != nil {
			panic(err)
		}
		defer result.Close()

		users := make([]client, 0)
		for result.Next() {
			var id int
			var role string
			var gender string
			var diagnosis int
			err = result.Scan(&id, &role, &gender, &diagnosis)
			if err != nil {
				panic(err)
			}

			users = append(users, client{id, role, gender, diagnosis})
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
		// Parse JSON from body of request

		// deets from JSON
		role := ""
		gender := ""
		diagnosis := ""

		sqlStatement := `INSERT INTO users(id,role,gender,diagnosis)
		VALUES(DEFAULT,$1,$2,$3);`

		_, err := db.Exec(sqlStatement, role, gender, diagnosis)
		if err != nil {
			// TODO - dont panic
			panic(err)
		}

		// Appropriate http response
	}
}
