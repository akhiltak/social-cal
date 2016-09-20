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
	"io/ioutil"
	"time"
)

type Friend struct {
	fname       string
	mname       string
	lname       string
	nick        string
	loc         Location
	birthday    time.Time
	anniversary time.Time
}

func (f *Friend) GetName() string {
	return f.fname + f.mname + f.lname
}

func (f *Friend) SetName(first, middle, last string) {
	f.fname = first
	f.mname = middle
	f.lname = last
}

func (f *Friend) GetNick() string {
	return f.nick
}

func (f *Friend) SetNick(n string) {
	f.nick = n
}

func (f *Friend) GetLocation() Location {
	return f.loc
}

func (f *Friend) GetBirthday() time.Time {
	return f.birthday
}

func (f *Friend) GetAnniversary() time.Time {
	return f.anniversary
}

func (f *Friend) Save() error {
	return ioutil.WriteFile(f.fname+".txt", []byte(f.fname), 0600)
}

type Location struct {
	country string
	city    string
	pincode string
}

func LoadFriends(fname string) (*Friend, error) {
	text, err := ioutil.ReadFile(fname + ".txt")
	if err != nil {
		return nil, err
	}
	f := Friend{nick: string(text)}
	return &f, nil
}
