package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/r3labs/diff/v3"
	"github.com/wphylici/go-intensive/module01/ex01/src/models"
	"io"
	"os"
	"path/filepath"
)

type DBReader interface {
	Reader(f *os.File) models.Recipes
}

type jsons struct{}
type xmls struct{}

func (j jsons) Reader(f *os.File) models.Recipes {
	var database models.Recipes

	dataJson, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("error reading file: err:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(dataJson, &database)
	if err != nil {
		fmt.Println("error unmarshal data: err:", err)
		os.Exit(1)
	}
	return database
}

func (x xmls) Reader(f *os.File) models.Recipes {
	var database models.Recipes

	dataXml, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("error reading file: err:", err)
		os.Exit(1)
	}

	err = xml.Unmarshal(dataXml, &database)
	if err != nil {
		fmt.Println("error unmarshal data: err:", err)
		os.Exit(1)
	}
	return database
}

func ReaderFromFile(filename string) models.Recipes {
	var database models.Recipes
	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		switch filepath.Ext(filename) {
		case ".xml":
			database = x.Reader(f)
		case ".json":
			database = j.Reader(f)
		default:
			fmt.Println("error: invalid file format")
			os.Exit(1)
		}
	} else {
		fmt.Println("error: file not specified")
		os.Exit(1)
	}
	return database
}

func checkAndOutputChanges(oldRecipe models.Recipes, newRecipe models.Recipes) {
	changelog, _ := diff.Diff(oldRecipe.Cake, newRecipe.Cake, diff.DisableStructValues())
	for _, el := range changelog {
		if el.Type == diff.CREATE {
			if len(el.Path) == 1 {
				fmt.Printf("ADDED cake \"%s\"\n", el.Path[0])
			} else if el.Path[1] == "Ingredients" {
				fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", el.Path[2], el.Path[0])
			}
		}
		if el.Type == diff.UPDATE {
			if el.Path[1] == "Time" {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s min\" instead of \"%s min\"\n",
					el.Path[0], el.To, el.From)
			} else if el.Path[1] == "Ingredients" {
				if el.Path[3] == "Count" {
					fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
						el.Path[2], el.Path[0], el.To, el.From)
				} else if el.Path[3] == "Unit" {
					if el.From == "" {
						fmt.Printf("ADDED unit %s for ingredient \"%s\" for cake  \"%s\"\n",
							el.To, el.Path[2], el.Path[0])
					} else if el.To == "" {
						fmt.Printf("REMOVED unit %s for ingredient \"%s\" for cake  \"%s\"\n",
							el.To, el.Path[2], el.Path[0])
					} else {
						fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n",
							el.Path[2], el.Path[0], el.To, el.From)
					}
				}
			}
		}
		if el.Type == diff.DELETE {
			if len(el.Path) == 1 {
				fmt.Printf("REMOVED cake \"%s\"\n", el.Path[0])
			} else if el.Path[1] == "Ingredients" {
				fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", el.Path[2], el.Path[0])
			}
		}
	}
}

var j, x DBReader = jsons{}, xmls{}

func main() {
	var filenameOld string
	var filenameNew string

	flag.StringVar(&filenameOld, "old", "", "file name")
	flag.StringVar(&filenameNew, "new", "", "file name")
	flag.Parse()

	oldRecipe := ReaderFromFile(filenameOld)
	newRecipe := ReaderFromFile(filenameNew)

	checkAndOutputChanges(oldRecipe, newRecipe)
}
