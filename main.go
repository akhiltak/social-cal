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
	"github.com/gin-gonic/gin"
)

type temp struct {
}

func viewHandler(c *gin.Context) {
	// TODO: remove hard coding for filename, take as flag input
	fl, err := model.LoadFriends("friends")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, fl)
}

func addHandler(c *gin.Context) {
	f, err := model.AddFriend(c.Param("name"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fl := []model.Friend{f}
	c.JSON(http.StatusOK, fl[0])
	// renderTemplate(w, "../../asset/add", &fl[0])
}

func editHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not here yet."})
}

func githubWebhook(c *gin.Context) {
	t := temp{}
	err := json.NewDecoder(c.Request.Body).Decode(t)
	fmt.Println(c.Request.Body)
	fmt.Println(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, t)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, data)
}

func notFound(c *gin.Context) {
	c.String(http.StatusOK, "Everything is osum.")
	// renderTemplate(w, "../../asset/not_found", fmt.Errorf("%v", fmt.Errorf("%v", r.URL.Path)))
}

func main() {
	fmt.Println("Running Friends Calendar and Events application...")

	r := gin.Default()
	// define routes
	r.GET("/", notFound)
	r.GET("/friends", viewHandler)
	r.GET("/add/:name", addHandler)
	r.GET("/edit", editHandler)
	r.GET("/payload", githubWebhook)

	fmt.Println("Ready!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
