package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/janjitsu/gossaurus/internal/entities"
	"log"
	"os"
	"strings"
)

type CreateDictionaryParams struct {
	From        string
	To          string
	EntriesFile string
	Title       string
	Author      string
	Cover       string
}

func parseTSV(filename string) [][]string {
	fmt.Println(filename)
	csvFile, _ := os.Open(filename)
	r := csv.NewReader(csvFile)
	r.Comma = '\t'         // Set tab as delimiter
	r.FieldsPerRecord = -1 // Allow variable number of fields

	// Read all records
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read TSV %v reason: %v", filename, err)
		panic(err)
	}
	return records
}

func parseEntries(filename string) []entities.Entry {
	records := parseTSV(filename)
	//@TODO need a validator, tell me if some line is wrong and not crash the app
	var entries []entities.Entry
	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}
		//@TODO this is ugly
		var entry entities.Entry
		entry.Title = record[0]
		//@TODO if there are no inflexsions, still is setting an empty string
		entry.Inflexions = strings.Split(record[1], ",")
		entry.Grammar = record[2]
		entry.Meaning1 = record[3]
		entry.Meaning2 = record[4]
		entry.Meaning3 = record[5]
		entry.Meaning4 = record[6]
		entry.Example1 = record[7]
		entry.Example2 = record[8]
		entry.Example3 = record[9]
		entry.Note1 = record[10]
		entry.Note2 = record[11]
		entry.Declensions = record[12]

		entries = append(entries, entry)
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
