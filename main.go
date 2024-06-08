package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// global db
var db *sql.DB

// stroing secrets in something like vault/aws secrets manager and retrieving them on using aws api or vault api. Harcoding them here just to make them work.
const (
	JWT_SECRET = ";qP7y5hTPBRbq%Q9K|R6s0>H£0'9L8GU"
	DB_USER    = "root"
	DB_PORT    = "3303"
	DB_PASS    = ""
	DB_HOST    = "localhost"
	DB_NAME    = "services_sample"
)

func main() {
	//Creating DB Connection
	dbDriver := "mysql"
	var errdb error
	db, errdb = sql.Open(dbDriver, DB_USER+":"+DB_PASS+"@("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if errdb != nil {
		panic(errdb.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 15)

	r := mux.NewRouter()
	r.Use(corsMiddleware())
	r.Methods("OPTIONS")
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})
	r.HandleFunc("/users/login", usersLoginHandler).Methods("POST")
	r.HandleFunc("/orgs", authMiddleware(orgsListHandler)).Methods("GET")
	r.HandleFunc("/orgs/{orgId}/services", authMiddleware(servicesListHandler)).Methods("GET")
	r.HandleFunc("/orgs/services/{serviceId}", authMiddleware(servicesInfoHandler)).Methods("GET")

	log.Println("Starting up server on port 8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
}
