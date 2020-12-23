package userRepository

import (
	"database/sql"
	"fmt"
	"log"
	"user_managment/model"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			log.Fatal(err)
		}
	}
}

func (ur UserRepository) GetUsers(db *sql.DB) []model.User {
	log.Println("Loading user(s) from database.")
	users := []model.User{}
	rows, err := db.Query("SELECT * FROM users")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.MobileNumber, &user.Address, &user.City, &user.Role, &user.Status, &user.CreationDate)
		logFatal(err)
		users = append(users, user)
	}
	log.Printf("Retrieved %v user(s) successfully.", len(users))
	return users
}

func (ur UserRepository) GetUser(db *sql.DB, id int) model.User {
	log.Println("Get user with id : ", id)
	rows := db.QueryRow("SELECT * FROM users where user_id=$1", id)
	user := model.User{}
	err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.MobileNumber, &user.Address, &user.City, &user.Role, &user.Status, &user.CreationDate)
	logFatal(err)
	log.Println("User loaded successfully with id", user.UserId)
	return user
}

func (ur UserRepository) AddUser(db *sql.DB, user model.User) int {
	log.Println("User creating with name :", user.FirstName+" "+user.LastName)
	err := db.QueryRow("insert into users (first_name, last_name, email, mobile_number, address, city, role) values ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id;", user.FirstName, user.LastName, user.Email, user.MobileNumber, user.Address, user.City, user.Role).Scan(&user.UserId)
	logFatal(err)
	log.Println("User added successfully with id", user.UserId)
	return user.UserId
}

func (ur UserRepository) UpdateUser(db *sql.DB, user model.User) int {
	log.Println("User updating with id :", user.UserId)
	var UserId int
	row := db.QueryRow("update users set first_name=$1, last_name=$2, email=$3, mobile_number=$4, address=$5, city=$6, role=$7, status=$8 where user_id=$9 RETURNING user_id;", &user.FirstName, &user.LastName, &user.Email, &user.MobileNumber, &user.Address, &user.City, &user.Role, &user.Status, &user.UserId)
	err := row.Scan(&UserId)
	logFatal(err)
	log.Println("User updated successfully with id", UserId)
	return UserId
}

func (ur UserRepository) DeleteUser(db *sql.DB, id int) int {
	log.Println("User deleting with id", id)
	var UserId int
	row := db.QueryRow("delete from users where user_id=$1", id)
	err := row.Scan(&UserId)
	logFatal(err)
	log.Println("User deleted successfully with id", id)
	return id
}
