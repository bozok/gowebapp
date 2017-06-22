package main

import (
	"encoding/json"
	"fmt"
	"gowebapp/core"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/nu7hatch/gouuid"
)

// Config file content
type Config struct {
	WebServerIP   string `json:"WebServerIP"`
	WebServerPort string `json:"WebServerPort"`
}

var cfg Config

func init() {
	loadConfig("config.json")
	core.Store = sessions.NewCookieStore([]byte("tasimasudanmatematikprofersoruolmaz"))
	core.Store.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 60, // 1 minute
	}
}

func loadConfig(cfgFileName string) {
	// first, load the JSON file and parse it into /cfg/
	f, err := os.Open(cfgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	jsonParser := json.NewDecoder(f)
	if err = jsonParser.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
}

func main() {
	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/", core.ServeContent)
	gorillaRoute.HandleFunc("/{page_alias}", core.ServeContent)

	http.HandleFunc("/img/", core.ServeResources)
	http.HandleFunc("/css/", core.ServeResources)
	http.HandleFunc("/js/", core.ServeResources)

	http.Handle("/", gorillaRoute)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/SignIn", signinAttempt)
	http.HandleFunc("/SignOut", signoutAttempt)

	http.ListenAndServe(":"+cfg.WebServerPort, nil)
}

func signinAttempt(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	_user := r.FormValue("user")
	_password := r.FormValue("pass")
	if _user == "faydin" && _password == "123qwe" {
		session, _ := core.Store.Get(r, "gwa-logged-in")
		id, _ := uuid.NewV4()
		session.Values["sid"] = id.String()
		session.Values["uid"] = 17
		session.Values["fullname"] = "Ali Fatih AYDIN"
		session.Options.MaxAge = 60
		session.Save(r, w)
		core.ResponseJSONRow(w, r, 7, "singin success!", "", "", true)
	} else {
		core.ResponseJSONRow(w, r, 7, "signin fail!", "", "", false)
	}
}

func signoutAttempt(w http.ResponseWriter, r *http.Request) {
	session, err := core.Store.Get(r, "gwa-logged-in")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		session.Options.MaxAge = -1
		session.Save(r, w)
	}
	core.ResponseJSONRow(w, r, 7, "signout ok!", "", "", true)
}
