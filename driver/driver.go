package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB
var drop_roles string = "DROP TABLE IF EXISTS roles"
var drop_users string = "DROP TABLE IF EXISTS users"

var create_roles string = "CREATE TABLE roles (role_id serial PRIMARY KEY, role_name VARCHAR (255) NOT NULL)"
var create_users string = "CREATE TABLE users (user_id serial PRIMARY KEY, first_name VARCHAR (255) NOT NULL, last_name VARCHAR (255) NOT NULL,	email VARCHAR (255) NOT NULL, mobile_number serial NOT NULL, address TEXT, city VARCHAR (255) NOT NULL,	role VARCHAR (255) NOT NULL, status VARCHAR (255) DEFAULT 'Active',	created_on DATE NOT NULL DEFAULT CURRENT_DATE)"

var insert_roles string = "INSERT INTO roles (role_name) VALUES ($1)"

//INSERT INTO roles (role_name) VALUES ('Super Admin');

var insert_users string = "INSERT INTO users (first_name, last_name, email, mobile_number, address, city,role) VALUES ($1, $2, $3, $4, $5, $6, $7)"

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateConnection() *sql.DB {
	if db != nil {
		return db
	}
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err := sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)
	return db
}

func CreateTables(db *sql.DB) {
	_, err := db.Exec(drop_roles)
	logFatal(err)

	_, err = db.Exec(drop_users)
	logFatal(err)

	_, err = db.Exec(create_roles)
	logFatal(err)

	_, err = db.Exec(create_users)
	logFatal(err)
}

func InsertData(db *sql.DB) {
	stmt, err := db.Prepare(insert_roles)
	logFatal(err)
	_, err = stmt.Exec("Super Admin")
	logFatal(err)
	_, err = stmt.Exec("Admin")
	logFatal(err)
	_, err = stmt.Exec("Engineer")
	logFatal(err)

	stmt, err = db.Prepare(insert_users)
	logFatal(err)
	_, err = stmt.Exec("Malyadri", "Patibandla", "malya4java@gmail.com", "959595", "Chennai", "Chennai", "Admin")
	logFatal(err)
}
