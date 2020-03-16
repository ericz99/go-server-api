package routers

import (
	"go-server-api/app/controllers"
	"go-server-api/app/middlewares"
	"go-server-api/app/utils"
	"log"
	"net/http"

	// # mysql lib
	// _ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App (Model)
type App struct {
	Router *mux.Router
}

var db = utils.GetDB()

// Initialize Server Stuff
func (a *App) Initialize() {
	// # setup router
	a.Router = mux.NewRouter()
	a.SetupRouters()
	// # apply any middleware
	a.ApplyMiddleware(middlewares.LoggerMiddleware)
	// # add cors method for api
	// a.ApplyMiddleware(mux.CORSMethodMiddleware(a.Router))
}

// ApplyMiddleware METHOD
func (a *App) ApplyMiddleware(middleware func(next http.Handler) http.Handler) {
	a.Router.Use(middleware)
}

// SetupRouters METHOD
func (a *App) SetupRouters() {
	a.Get("/api/book/test", a.getTestRoute)
	a.Post("/api/book/save", a.saveBookRoute)
	a.Get("/api/book/{id}", a.getBookRoute)
	a.Delete("/api/book/{id}", a.deleteBookRoute)
	a.Get("/api/books", a.getAllBookRoute)
}

// Get METHOD
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post METHOD
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Delete METHOD
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Put METHOD
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) getTestRoute(w http.ResponseWriter, r *http.Request) {
	controllers.GetBookTestRoute(w, r)
}

func (a *App) saveBookRoute(w http.ResponseWriter, r *http.Request) {
	controllers.SaveBookRoute(db, w, r)
}

func (a *App) getBookRoute(w http.ResponseWriter, r *http.Request) {
	controllers.GetBookRoute(db, w, r)
}

func (a *App) deleteBookRoute(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteBookRoute(db, w, r)
}

func (a *App) getAllBookRoute(w http.ResponseWriter, r *http.Request) {
	controllers.GetAllBookRoute(db, w, r)
}

// Run Server
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
