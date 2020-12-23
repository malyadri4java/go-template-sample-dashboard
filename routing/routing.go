package routing

import (
	"database/sql"
	"net/http"
	"user_managment/controller"
	"user_managment/driver"
	"user_managment/utils"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func Router() *mux.Router {
	utils.LoadTemplates("templates/layouts/*.html")

	db = driver.CreateConnection()
	driver.CreateTables(db)
	driver.InsertData(db)

	control := controller.Controller{}
	contForm := controller.ControllerForm{}

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/users/add", AddUserHandler).Methods("GET")
	router.HandleFunc("/users/edit/{id}", contForm.EditUserForm(db)).Methods("GET")
	router.HandleFunc("/users/delete/{id}", control.DeleteUser(db, "/users")).Methods("POST")
	router.HandleFunc("/users", contForm.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", control.UpdateUser(db, "/users")).Methods("POST")
	router.HandleFunc("/users", control.AddUser(db, "/users")).Methods("POST")

	// old user mangment UI.
	router.HandleFunc("/ui/users", contForm.GetUsersV1(db)).Methods("GET")
	router.HandleFunc("/ui/users/{id}", control.UpdateUser(db, "/ui/users")).Methods("POST")
	router.HandleFunc("/ui/users", control.AddUser(db, "/ui/users")).Methods("POST")
	router.HandleFunc("/ui/users/delete/{id}", control.DeleteUser(db, "/ui/users")).Methods("POST")

	// API Access
	router.HandleFunc("/api/users", control.GetUsers(db)).Methods("GET")
	router.HandleFunc("/api/users/{id}", control.GetUser(db)).Methods("GET")
	return router
}

func StaticHomeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/index.html")
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	utils.ExecuteTemplate(w, "index.html", nil) // here nil means model object is null
}

func AddUserHandler(w http.ResponseWriter, req *http.Request) {
	utils.ExecuteTemplate(w, "user_add.html", nil)
}
