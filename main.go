package main

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var themeName = getThemeName()
var staticPages = populateStaticPages()

func main() {
	serveWeb()
}

func serveWeb() {
	gorillaRoute := mux.NewRouter()

	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{page_alias}", serveContent)

	http.HandleFunc("/img/", serveResources)
	http.HandleFunc("/css/", serveResources)
	http.HandleFunc("/js/", serveResources)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)

}

func serveContent(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]
	if page_alias == "" {
		page_alias = "home"
	}

	staticPage := staticPages.Lookup(page_alias + ".html")
	if staticPage == nil {
		staticPage = staticPages.Lookup("404.html")
		w.WriteHeader(404)
	}
	//w.Write([]byte("Hello Universe! " + page_alias))
	staticPage.Execute(w, nil)
}

func getThemeName() string {
	return "bs4"
}

// Retrieve all files under given folder and its subsequent folders
func populateStaticPages() *template.Template {
	result := template.New("pages")
	basePath := "pages"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePaths := new([]string)
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		// if !pathInfo.IsDir() {
		// 	*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		// }
	}
	result.ParseFiles(*templatePaths...)
	return result
}

// serve resources of types js, img, css files
func serveResources(w http.ResponseWriter, req *http.Request) {
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

	log.Println(path)
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
