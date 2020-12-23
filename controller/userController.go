package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"user_managment/model"
	userRepository "user_managment/repository"

	"github.com/gorilla/mux"
)

type Controller struct{}

var repo = userRepository.UserRepository{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) AddUser(db *sql.DB, rUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		mobileNumber, err := strconv.Atoi(req.FormValue("MobileNumber"))
		logFatal(err)

		user := model.User{
			FirstName:    req.FormValue("FirstName"),
			LastName:     req.FormValue("LastName"),
			Email:        req.FormValue("Email"),
			MobileNumber: mobileNumber,
			Address:      req.FormValue("Address"),
			City:         req.FormValue("City"),
			Role:         req.FormValue("Role"),
			Status:       req.FormValue("Status"),
		}
		repo.AddUser(db, user)
		http.Redirect(w, req, rUrl, 302)
	}
}

func (c Controller) UpdateUser(db *sql.DB, rUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		userId, err := strconv.Atoi(params["id"])
		logFatal(err)

		req.ParseForm()
		mobileNumber, err := strconv.Atoi(req.FormValue("MobileNumber"))
		logFatal(err)

		user := model.User{
			UserId:       userId,
			FirstName:    req.FormValue("FirstName"),
			LastName:     req.FormValue("LastName"),
			Email:        req.FormValue("Email"),
			MobileNumber: mobileNumber,
			Address:      req.FormValue("Address"),
			City:         req.FormValue("City"),
			Role:         req.FormValue("Role"),
			Status:       req.FormValue("Status"),
		}
		repo.UpdateUser(db, user)
		http.Redirect(w, req, rUrl, 302)
	}
}

func (c Controller) DeleteUser(db *sql.DB, rUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		userId, err := strconv.Atoi(params["id"])
		logFatal(err)
		repo.DeleteUser(db, userId)
		http.Redirect(w, req, rUrl, 302)
	}
}

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users := []model.User{}
		users = repo.GetUsers(db)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func (c Controller) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := model.User{}
		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		user = repo.GetUser(db, id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
