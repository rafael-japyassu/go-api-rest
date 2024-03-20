package models

type Personality struct {
	Name    string `json:"name"`
	History string `json:"history"`
}

var Personalities []Personality
