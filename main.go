package main

import (
	"fmt"
	"github.com/dmmmd/scrumpoker/controller"
	"github.com/dmmmd/scrumpoker/grooming_session"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	r := buildRouter()

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func buildRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", actionHomepage)
	r.Get("/actionEcho", actionEcho)

	r.Route("/grooming_sessions", grooming_session.Router)

	return r
}

// Actions

func actionHomepage(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	body := "<html>" +
		"<head>" +
		"<title>Who has time for this?</title>" +
		"</head>" +
		"<body>" +
		"<ul>" +
		"<li><a href=\"./grooming_sessions\">Grooming Sessions</a></li>" +
		"<li><a href=\"./actionEcho?foo=bar&baz=umad?&debug=true\">Echo</a></li>" +
		"</ul>" +
		"</body>" +
		"</html>"
	controller.SendRawResponse(w, body)
}

func actionEcho(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	printDebug(r.URL.Path, r.Form)
	body := ""
	for k, v := range r.Form {
		body += fmt.Sprintf("%s:\t%s\n", k, strings.Join(v, ""))
	}
	controller.SendRawResponse(w, body)
}

// Other

func printDebug(path string, query url.Values) {
	if query.Get("debug") != "" {
		fmt.Println(query)
		fmt.Println("path", path)
		for k, v := range query {
			fmt.Println(fmt.Sprintf("%s\t=>\t%s", k, v))
		}
	}
}
