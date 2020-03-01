package model

import (
"encoding/json"
)

type Human struct {
	FirstName  string ` json:"first_name" xml:"first_name" `
	SecondName string `json:"second_name" xml:"second_name" `
	Age        string  `json:"age" xml:"age"`
}
func (h Human) ForSendJSON() string {

	humanStr, _ :=json.Marshal(h)
	return string(humanStr)
}
type boss interface {
	bossPosition()
}
type worker interface {
	getProfession() string
	getPosition() string
	getName() string
}
type doctor struct {
	Human
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver doctor) getProfession() string {
	return resiver.Profession
}
func (resiver doctor) getPosition() string {
	return resiver.Position
}
func (resiver doctor) getName() string {
	return resiver.FirstName
}

type manager struct {
	Human
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver manager) getProfession() string {
	return resiver.Profession
}
func (resiver manager) getPosition() string {
	return resiver.Position
}
func (resiver manager) getName() string {
	return resiver.FirstName
}

type sportsmen struct {
	Human
	Sport      string `json:"sport" xml:"sport"`
	Score      string `json:"score" xml:"score"`
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver sportsmen) getProfession() string {
	return resiver.Profession
}
func (resiver sportsmen) getPosition() string {
	return resiver.Position
}
func (resiver sportsmen) getName() string {
	return resiver.FirstName
}

type programmer struct {
	Human
	Profession string `json:"specialty" xml:"specialty"`
	Position   string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver programmer) getProfession() string {
	return resiver.Profession
}
func (resiver programmer) getPosition() string {
	return resiver.Position
}
func (resiver programmer) getName() string {
	return resiver.FirstName
}

type teacher struct {
	Human
	Field            string `json:"field" xml:"field"`
	NumberOfStudents int    `json:"number_of_students" xml:"number_of_students"`
	Profession       string `json:"specialty" xml:"specialty"`
	Position         string `json:"list_of_patients,omitempty" xml:"list_of_patients"`
}

func (resiver teacher) getProfession() string {
	return resiver.Profession
}
func (resiver teacher) getPosition() string {
	return resiver.Position
}
func (resiver teacher) getName() string {
	return resiver.FirstName
}
