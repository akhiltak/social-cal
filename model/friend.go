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

package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Friend defines structure for yaari-dosti
type Friend struct {
	Fname       string    `json:"first_name"`
	Mname       string    `json:"middle_name,omitempty"`
	Lname       string    `json:"last_name,omitempty"`
	Nick        string    `json:"nick_name,omitempty"`
	Loc         Location  `json:"location,omitempty"`
	Birthday    time.Time `json:"birthday,omitempty"`
	Anniversary time.Time `json:"anniversary,omitempty"`
}

// GetName is getter method for entire name
func (f *Friend) GetName() string {
	return strings.TrimSpace(f.Fname+" "+f.Mname) + " " + f.Lname
}

// SetName is setter method for fields first_name, middle_name and last_name
func (f *Friend) SetName(first, middle, last string) {
	f.Fname = first
	f.Mname = middle
	f.Lname = last
}

// GetBirthday is getter method for field birthday
func (f *Friend) GetBirthday() time.Time {
	return f.Birthday
}

// GetAnniversary is getter method for field anniversary
func (f *Friend) GetAnniversary() time.Time {
	return f.Anniversary
}

// save saves a friend to friend list (file or database)
func (f *Friend) save() error {
	file, err := os.OpenFile("friends.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString("\n" + f.GetName())
	return err
}

// AddFriend adds a new friend
func AddFriend(n string) (Friend, error) {
	f := Friend{
		Fname: n,
		Loc: Location{
			City:    "Mumbai",
			Country: "India",
		},
	}
	return f, f.save()
}

// LoadFriends gets all friends from data storage
func LoadFriends(filename string) ([]Friend, error) {
	file, err := ioutil.ReadFile(filename + ".json")
	if err != nil {
		return nil, err
	}
	var fl []Friend
	if err = json.Unmarshal(file, &fl); err != nil {
		return nil, err
	}
	return fl, nil
}

// Location defines struct to store locations
type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Pincode string `json:"pincode,omitempty"`
}

func (l *Location) String() string {
	return l.City + ", " + l.Country + "[" + l.Pincode + "]"
}
