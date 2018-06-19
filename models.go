package main

import "reflect"

type SoloData struct {
	Key  string			`msgpack:"key" json:"key" binding:"required"`
	Text string 		`msgpack:"text" json:"text" binding:"required"`
}

type MultiData struct {
	Key  string			`msgpack:"key" binding:"required"`
	Texts []string 		`msgpack:"texts" binding:"required"`
}

func (data *SoloData) validate() {
	if reflect.ValueOf(data.Key).Len() != 24 {
		panic("Key length should be 24 letters.")
	}
}

func (data *MultiData) validate() {
	if reflect.ValueOf(data.Key).Len() != 24 {
		panic("Key length should be 24 letters.")
	}
}

