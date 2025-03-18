package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/janjitsu/gossaurus/internal/entities"
	"io"
	"log"
	"os"
)

type CreateDictionaryParams struct {
	From        string
	To          string
	EntriesFile string
	Title       string
	Author      string
	Cover       string
}

func parseEntries(filename string) []entities.Entry {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to read TSV %v reason: %v", filename, err)
		panic(err)
	}
	defer csvFile.Close()
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = '\t'
		r.FieldsPerRecord = -1 // Allow variable number of fields
		return r               // Allows use dot as delimiter and use quotes in CSV
	})

	var entries []entities.Entry
	if err := gocsv.UnmarshalFile(csvFile, &entries); err != nil {
		panic(err)
	}

	return entries
}

func ParseDictionary(c CreateDictionaryParams) entities.Dictionary {
	entries := parseEntries(c.EntriesFile)
	fmt.Println(entries)

	//create a page
	var pages []entities.Page
	pages = append(pages, entities.Page{"page_1", entries})
	//create a dictionary
	version := "version1"
	dictionary := entities.Dictionary{
		version,
		c.Title,
		c.From,
		c.To,
		c.Author,
		c.Cover,
		pages,
	}

	return dictionary
}
