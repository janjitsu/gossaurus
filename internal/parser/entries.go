package parser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/janjitsu/gossaurus/internal/entities"
	"io"
	"log"
	"os"
	"path/filepath"
)

// @TODO make it an interface and inject parser types
func parseEntriesFromCSV(filename string) (entries []entities.Entry, err error) {
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

	if err := gocsv.UnmarshalFile(csvFile, &entries); err != nil {
		panic(err)
	}

	return entries, nil
}

func parseEntriesFromJSON(filename string) (entries []entities.Entry, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = json.Unmarshal(byteValue, &entries)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	// Print the parsed data
	fmt.Println("Parsed JSON:", entries)
	return entries, nil
}

func parseEntries(filename string) (entries []entities.Entry, err error) {
	ext := filepath.Ext(filename)
	if ext == ".json" {
		return parseEntriesFromJSON(filename)
	} else if ext == ".tsv" {
		return parseEntriesFromCSV(filename)
	} else {
		return nil, fmt.Errorf("unsupported file (%s), use json or tsv", filename)
	}
}
