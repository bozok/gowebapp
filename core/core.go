package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

type jsonreturn struct {
	Code    int
	Message string
	Page    string
	Data    string
	Ok      bool
}

// ResponseJSONRow ...
func ResponseJSONRow(w http.ResponseWriter, r *http.Request, code int, message string, page string, data string, ok bool) {
	w.Header().Set("Content-Type", "application/json")
	jsonfinal := jsonreturn{code, message, page, data, ok}
	js, _ := json.Marshal(jsonfinal)
	w.Write(js)
}

var themeName = getThemeName()

func getThemeName() string {
	return "bs4"
}

// ServeResources serve resources of types js, img, css files
func ServeResources(w http.ResponseWriter, req *http.Request) {
	path := "public/" + themeName + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}
	//log.Println(path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func ServeContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]
	sc := checkSession(w, r)
	if sc == 0 {
		page_alias = "signin"
	} else {
		if page_alias == "" || page_alias == "signin" {
			page_alias = "home"
		}
	}
	staticPage := staticPages.Lookup(page_alias + ".html")
	if staticPage == nil {
		staticPage = staticPages.Lookup("404.html")
		w.WriteHeader(404)
	}
	// Values to pass into the template
	context := defaultContext{}
	context.Title = page_alias
	context.ErrorMsg = ""
	context.SuccessMsg = ""
	session, err := Store.Get(r, "gwa-logged-in")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		if session.IsNew {

		} else {
			context.UserFullName = session.Values["fullname"].(string)
			context.UserID = session.Values["uid"].(int)
		}
	}
	//staticPage.Execute(w, nil)
	staticPage.Execute(w, context)
}

// Struct to pass into the template
type defaultContext struct {
	Title        string
	ErrorMsg     string
	SuccessMsg   string
	UserFullName string
	UserID       int
}

var staticPages = populateStaticPages()

// Retrieve all files under given folder and its subsequent folders
func populateStaticPages() *template.Template {
	result := template.New("pages")
	templatePaths := new([]string)

	basePath := "pages"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		//log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		// if !pathInfo.IsDir() {
		// 	*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		// }
	}

	basePath = "themes"
	templateFolder, _ = os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ = templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		//log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
	}
	result.ParseFiles(*templatePaths...)
	return result
}

func checkSession(w http.ResponseWriter, r *http.Request) byte {
	var retval byte
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	session, err := Store.Get(r, "gwa-logged-in")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		if session.IsNew {
			session.Options.MaxAge = -1
			retval = 0
		} else {
			session.Options.MaxAge = 60
			retval = 1
		}
	}
	session.Save(r, w)
	return retval
}
