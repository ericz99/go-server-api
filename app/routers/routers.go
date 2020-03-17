package routers

import (
	AuthController "go-server-api/app/auth"
	V1Controller "go-server-api/app/controllers/api/v1"
	V2Controller "go-server-api/app/controllers/api/v2"
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

// # get db instance
var db = utils.GetDB()

// Initialize Server Stuff
func (a *App) Initialize() {
	// # setup router
	a.Router = mux.NewRouter()
	a.SetupAuthAPIRouter()
	a.SetupAPIRouter("/api/v1")
	a.SetupAPIRouter("/api/v2")
	// # apply any middleware
	a.ApplyMiddleware(middlewares.LoggerMiddleware)
	a.ApplyMiddleware(middlewares.CommonMiddleware)
	// a.ApplyMiddleware(middlewares.AuthMiddleware)
	// # add cors method for api
	// a.ApplyMiddleware(mux.CORSMethodMiddleware(a.Router))
}

// ApplyMiddleware METHOD
func (a *App) ApplyMiddleware(middleware func(next http.Handler) http.Handler) {
	a.Router.Use(middleware)
}

// SetupAPIRouter METHOD
func (a *App) SetupAPIRouter(path string) {
	// # creates a subrouter path
	s := a.Router.PathPrefix(path).Subrouter()
	s.Use(middlewares.AuthMiddleware)

	// # here we setup our api router path version
	if path == "/api/v1" {
		// # setup router for API version 1
		a.SetupAPIV1Router(s)
	} else {
		// # setup router for API version 1
		a.SetupAPIV2Router(s)
	}
}

// SetupAPIV1Router METHOD (VERSION API 1)
func (a *App) SetupAPIV1Router(s *mux.Router) {
	a.Get(s, "/status", a.statusBookRouteV1)
	a.Post(s, "/book/save", a.saveBookRouteV1)
	a.Get(s, "/book/{id}", a.getBookRouteV1)
	a.Delete(s, "/book/{id}", a.deleteBookRouteV1)
	a.Get(s, "/books", a.getAllBookRouteV1)
}

// SetupAPIV2Router METHOD (VERSION API 2)
func (a *App) SetupAPIV2Router(s *mux.Router) {
	a.Get(s, "/status", a.statusBookRouteV2)
	a.Post(s, "/book/save", a.saveBookRouteV2)
	a.Get(s, "/book/{id}", a.getBookRouteV2)
	a.Delete(s, "/book/{id}", a.deleteBookRouteV2)
	a.Get(s, "/books", a.getAllBookRouteV2)
}

//SetupAuthAPIRouter METHOD (SETUP AUTH HANDLER)
func (a *App) SetupAuthAPIRouter() {
	a.Post(a.Router, "/api/auth/login", a.loginRouteAPI)
	a.Post(a.Router, "/api/auth/register", a.registerRouteAPI)
}

// Get METHOD
func (a *App) Get(s *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.HandleFunc(path, f).Methods("GET")
}

// Post METHOD
func (a *App) Post(s *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.HandleFunc(path, f).Methods("POST")
}

// Delete METHOD
func (a *App) Delete(s *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.HandleFunc(path, f).Methods("DELETE")
}

// Put METHOD
func (a *App) Put(s *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.HandleFunc(path, f).Methods("PUT")
}

func (a *App) statusBookRouteV1(w http.ResponseWriter, r *http.Request) {
	V1Controller.GetBookStatusRoute(w, r)
}

func (a *App) saveBookRouteV1(w http.ResponseWriter, r *http.Request) {
	V1Controller.SaveBookRoute(db, w, r)
}

func (a *App) getBookRouteV1(w http.ResponseWriter, r *http.Request) {
	V1Controller.GetBookRoute(db, w, r)
}

func (a *App) deleteBookRouteV1(w http.ResponseWriter, r *http.Request) {
	V1Controller.DeleteBookRoute(db, w, r)
}

func (a *App) getAllBookRouteV1(w http.ResponseWriter, r *http.Request) {
	V1Controller.GetAllBookRoute(db, w, r)
}

func (a *App) statusBookRouteV2(w http.ResponseWriter, r *http.Request) {
	V2Controller.GetBookStatusRoute(w, r)
}

func (a *App) saveBookRouteV2(w http.ResponseWriter, r *http.Request) {
	V2Controller.SaveBookRoute(db, w, r)
}

func (a *App) getBookRouteV2(w http.ResponseWriter, r *http.Request) {
	V2Controller.GetBookRoute(db, w, r)
}

func (a *App) deleteBookRouteV2(w http.ResponseWriter, r *http.Request) {
	V2Controller.DeleteBookRoute(db, w, r)
}

func (a *App) getAllBookRouteV2(w http.ResponseWriter, r *http.Request) {
	V2Controller.GetAllBookRoute(db, w, r)
}

func (a *App) loginRouteAPI(w http.ResponseWriter, r *http.Request) {
	AuthController.Login(db, w, r)
}

func (a *App) registerRouteAPI(w http.ResponseWriter, r *http.Request) {
	AuthController.Register(db, w, r)
}

// Run Server
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
