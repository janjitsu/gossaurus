package parser

import (
	"fmt"
	"github.com/janjitsu/gossaurus/internal/entities"
)

type CreateDictionaryParams struct {
	From        string
	To          string
	EntriesFile string
	Title       string
	Author      string
	Cover       string
}

func ParseDictionary(c CreateDictionaryParams) entities.Dictionary {
	entries, _ := parseEntries(c.EntriesFile)
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
