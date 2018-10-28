/*
 * File: acheteur.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 23rd October 2018 5:45:57 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package model

import (
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// Acheteur is composed:
type Acheteur struct {
<<<<<<< 5bdf2231d118196baa3442a5639ee97ff02c5476
	Civilite    string    	`json:"Civilite"`
	Nom       	string    	`json:"Nom"`
	Prenom      string    	`json:"Prenom"`
	Age   		int			`json:"Age, float64"`
	Email 		string    	`json:"Email"`
=======
	Spectateur
	Email  string `json:"Email"`
	mapped map[string]interface{}
>>>>>>> feat(iserial): we dont pass by json marshall annotation, we convert struct to map, this allow us to have a universal way to pass data to IResponseWriter instance
}

// ToString return string conversion of marshal user
// absorb error...
func (a *Acheteur) ToString() string {
	ret, _ := a.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (a *Acheteur) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(a.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, a)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(a.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// Marshal implement ISerial
func (a Acheteur) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func (a Acheteur) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(a).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	a.mapped = mapped
	return mapped
}

func (a Acheteur) GetMapped() map[string]interface{} {
	if len(a.mapped) == 0 {
		a.toMap()
	}
	return a.mapped
}
