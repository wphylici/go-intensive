package models

import "encoding/xml"

type Recipes struct {
	XMLName xml.Name `json:"-" xml:"recipes" diff:"-"`
	Cake    []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	XMLName     xml.Name      `json:"-" xml:"cake" diff:"-"`
	Name        string        `json:"name" xml:"name" diff:"name,identifier"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Ingredients struct {
	XMLName xml.Name `json:"-" xml:"item" diff:"-"`
	Name    string   `json:"ingredient_name" xml:"itemname" diff:"ingredient_name,identifier"`
	Count   string   `json:"ingredient_count" xml:"itemcount"`
	Unit    string   `json:"ingredient_unit," xml:"itemunit,omitempty"`
}
