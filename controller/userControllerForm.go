package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"user_managment/model"
	"user_managment/utils"

	"github.com/gorilla/mux"
)

type ControllerForm struct{}

func (c ControllerForm) GetUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users := []model.User{}
		users = repo.GetUsers(db)
		utils.ExecuteTemplate(w, "user.html", users)
	}
}

func (c ControllerForm) GetUsersV1(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		users := []model.User{}
		users = repo.GetUsers(db)
		utils.ExecuteTemplateWitoutlayout(w, "old_user.html", users)
	}
}

func (c ControllerForm) EditUserForm(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := model.User{}
		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])

		user = repo.GetUser(db, id)
		utils.ExecuteTemplate(w, "profile.html", user)
	}
}
