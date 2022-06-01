package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/wphylici/go-intensive/module01/ex00/src/models"
	"io"
	"os"
	"path/filepath"
)

type DBReader interface {
	Reader(f *os.File) models.Recipes
}

type stolen struct{}
type original struct{}

func (s stolen) Reader(f *os.File) models.Recipes {
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

func (o original) Reader(f *os.File) models.Recipes {
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

func main() {
	var filename string
	var s, o DBReader = &stolen{}, &original{}

	flag.StringVar(&filename, "f", "", "file name")
	flag.Parse()

	if filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		switch filepath.Ext(filename) {
		case ".xml":
			database := o.Reader(f)
			dataJson, err := json.MarshalIndent(database, "", "    ")
			if err != nil {
				fmt.Println("error marshal data: err:", err)
				os.Exit(1)
			}
			fmt.Println(string(dataJson))
		case ".json":
			database := s.Reader(f)
			dataXml, err := xml.MarshalIndent(database, "", "    ")
			if err != nil {
				fmt.Println("error marshal data: err:", err)
				os.Exit(1)
			}
			fmt.Println(string(dataXml))
		default:
			fmt.Println("error: invalid file format")
			os.Exit(1)
		}
	} else {
		fmt.Println("error: file not specified")
		os.Exit(1)
	}
}
