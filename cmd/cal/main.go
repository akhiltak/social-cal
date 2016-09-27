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
	"fmt"
	"net/http"
	"strings"

	"github.com/akhiltak/social-cal/model"
)

func handler(w http.ResponseWriter, r *http.Request) {
	route := strings.Split(r.URL.Path[1:], "/")
	fmt.Fprintln(w, "<html><br />")
	switch route[0] {
	case "friends":
		fmt.Fprint(w, fetch())
	case "add":
		fmt.Fprint(w, add(route[1]))
	default:
		fmt.Fprintf(w, "ROUTING ERROR: Unknown route => <b>%v</b>\n", r.URL.Path)
	}
	fmt.Fprintln(w, "<br /></html>")
}

func add(n string) string {
	if err := model.AddFriend(n); err != nil {
		return fmt.Sprintf("Error adding friend:%v\n", err)
	}
	return fmt.Sprintf("Added:%v to friend list.", n)
}

func fetch() string {
	// TODO: remove hard coding for filename, take as flag input
	r, err := model.LoadFriends("friends")
	if err != nil {
		return fmt.Sprintf("Error fetching values:%v", err)
	}
	names := ""
	for _, n := range r {
		names = names + " " + n.GetNick()
	}
	return fmt.Sprintf("Names are:%v", names)
}

func main() {
	// define routes
	fmt.Println("Running Friends Calendar and Events application...")
	http.HandleFunc("/", handler)
	fmt.Println("Ready!")
	http.ListenAndServe(":8080", nil)
}
