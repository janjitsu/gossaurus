package command

import (
	"flag"
	"fmt"
	"github.com/janjitsu/gossaurus/internal/parser"
	"log"
	"os"
	"os/exec"

	"github.com/janjitsu/gossaurus/internal/builder/mobi"
)

type Command struct {
	From        string
	To          string
	EntriesFile string
	Title       string
	Author      string
	Cover       string
}

func BuildWithArgs() Command {
	cmd := Command{}
	flag.StringVar(&cmd.From, "from", "",
		"ISO two letter code of Origin Language. Ex. en")
	flag.StringVar(&cmd.To, "to", "",
		"ISO two letter code of Target Language. Ex. es ")
	flag.StringVar(&cmd.EntriesFile, "entries", "",
		"tsv file containing entries")

	// optional fields
	flag.StringVar(&cmd.Title, "title", "Converted Ebook",
		"Title of the Dictionary")
	flag.StringVar(&cmd.Author, "author", "jan.art.br", "Name of author")
	flag.StringVar(&cmd.Cover, "cover", "default-cover.jpg", "image for cover")
	flag.Parse()

	fmt.Println("Command", cmd)

	if cmd.From == "" {
		log.Fatalf("flag -from is required, received: %v", cmd.From)
	}

	if cmd.To == "" {
		log.Fatalf("flag -to is required, received: %v", cmd.To)
	}

	return cmd
}

func (c *Command) Execute() {
	dictionary := parser.ParseDictionary(parser.CreateDictionaryParams{
		c.From,
		c.To,
		c.EntriesFile,
		c.Title,
		c.Author,
		c.Cover,
	})
	// if version directory doesn't exist, create it
	os.MkdirAll(fmt.Sprintf("out/%s", dictionary.Id), os.ModePerm)

	//make entries into html page template
	for _, page := range dictionary.Pages {
		mobi.BuildPage(page, dictionary.Id)
	}
	//make dictionary into content template
	mobi.BuildDictionaryContent(dictionary)

	//save files to a zip package
	//ZipFolder("out/version1", "out/version1.zip")
	//fmt.Println("out/version1.zip")

	//@TODO call kindlegen
	kindlegen := "bin/kindlegen_linux/kindlegen"
	kindlegenArgs := "-c0"

	cmd := exec.Command(kindlegen, "out/version1/content.opf", kindlegenArgs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Errorf("kindlegen execution failed: %v", err)
	}
}
