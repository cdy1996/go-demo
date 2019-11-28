package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	_ "gopkg.in/go-playground/validator.v9"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	Middlewares *Middleware
	Config      *Env
}

//请求
type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

//响应
type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
}

// 初始化函数
func (a *App) Initialize(e *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Config = e
	a.Router = mux.NewRouter()
	a.Middlewares = &Middleware{}
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	m := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler)
	//a.Router.HandleFunc("/api/shorten", a.createShortlink).Methods("POST")
	//a.Router.HandleFunc("/api/info", a.getShortlinkInfo).Methods("GET")
	//a.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", a.redirect).Methods("GET")

	// 合并middleware
	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortlink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortlinkInfo)).Methods("GET")
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", m.ThenFunc(a.redirect)).Methods("GET")
}

func (a *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{
			http.StatusBadRequest,
			fmt.Errorf("parse parameters failed %v", r.Body)})
		return
	}
	if err := validator.Validate(req); err != nil {

		respondWithError(w, StatusError{
			http.StatusBadRequest,
			fmt.Errorf("validate parameters failed %v", req)})
		return
	}
	defer r.Body.Close()
	fmt.Printf("%v\n", req)
	s, e := a.Config.S.Shorten(req.URL, req.ExpirationInMinutes)
	if e != nil {
		respondWithError(w, e)
	} else {
		responseWithJson(w, http.StatusCreated, shortlinkResp{Shortlink: s})
	}

}

func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e)
		responseWithJson(w, e.Status(), e.Error())
	default:
		responseWithJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)

}

func (a *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")
	fmt.Printf("%s\n", s)

	//panic(s)
	info, e := a.Config.S.ShortlinkInfo(s)
	if e != nil {
		respondWithError(w, e)
	} else {
		responseWithJson(w, http.StatusOK, info)
	}

}
func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//fmt.Printf("%s\n", vars["shortlink"])

	u, e := a.Config.S.Unshorten(vars["shortlink"])
	if e != nil {
		respondWithError(w, e)
	} else {
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)

	}
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
