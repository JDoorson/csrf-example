package main

import (
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ShowMessageForm)
	r.HandleFunc("/post", PostMessageForm).Methods("POST")
	r.HandleFunc("/thanks", ShowThanksPage)

	http.ListenAndServe(":8080", csrf.Protect([]byte("secure-key-lol"), csrf.Secure(false))(r))
}

func ShowMessageForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

func PostMessageForm(w http.ResponseWriter, r *http.Request) {
	println("[POST]")
	println("URI: /post")
	println("DATA: " + r.FormValue("message"))
	println()

	http.Redirect(w, r, "/thanks?msg=" + r.FormValue("message"), http.StatusSeeOther)
}

func ShowThanksPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	t := template.Must(template.ParseFiles("thanks.html"))
	t.ExecuteTemplate(w, "thanks.html", map[string]interface{}{
		"message": query.Get("msg"),
	})
}