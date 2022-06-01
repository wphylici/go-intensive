package models

import "encoding/xml"

type Recipes struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cake    []Cake   `json:"cake" xml:"cake"`
}

type Cake struct {
	XMLName     xml.Name      `json:"-" xml:"cake"`
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Ingredients struct {
	XMLName xml.Name `json:"-" xml:"item"`
	Name    string   `json:"ingredient_name" xml:"itemname"`
	Count   string   `json:"ingredient_count" xml:"itemcount"`
	Unit    string   `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}
