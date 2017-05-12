/*
 * Copyright 2016 Akhil Tak (Stormblessed)
 *
 * This file is part of social-cal.

 * Socal-cal is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.

 * Social-cal is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.

 * You should have received a copy of the GNU General Public License
 * along with Foobar. If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/akhiltak/social-cal/model"
)

type temp struct {
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: remove hard coding for filename, take as flag input
	fl, err := model.LoadFriends("friends")
	if err != nil {
		renderTemplate(w, "../../asset/error", fmt.Errorf("Error fetching values:%v", err))
	}
	renderTemplate(w, "../../asset/friends", &fl)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	f, err := model.AddFriend(r.URL.Path[len("/add/"):])
	if err != nil {
		http.Redirect(w, r, "/friends", http.StatusFound)
		return
		// renderTemplate(w, "../../asset/error", err)
	}
	fl := []model.Friend{f}
	renderTemplate(w, "../../asset/add", &fl[0])
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// renderTemplate(w, "../../asset/edit", &f)
	renderTemplate(w, "../../asset/error", fmt.Errorf("%v", "Not implemented yet."))
}

func githubWebhook(w http.ResponseWriter, r *http.Request) {
	t := temp{}
	err := json.NewDecoder(r.Body).Decode(t)
	fmt.Println(r.Body)
	fmt.Println(t)
	if err != nil {
		renderTemplate(w, "../../asset/error", fmt.Errorf("Error fetching values:%v", err))
	}
	renderTemplate(w, "../../asset/default", t)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, data)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "../../asset/not_found", fmt.Errorf("%v", fmt.Errorf("%v", r.URL.Path)))
}

func main() {
	fmt.Println("Running Friends Calendar and Events application...")
	// define routes
	http.HandleFunc("/", notFound)
	http.HandleFunc("/friends", viewHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/payload", githubWebhook)

	fmt.Println("Ready!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
