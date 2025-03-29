package command

import (
	"fmt"
	"github.com/janjitsu/gossaurus/internal/parser"
	"os"
	"os/exec"

	"github.com/janjitsu/gossaurus/internal/builder/mobi"
)

type BuildCommand struct {
	From        string
	To          string
	EntriesFile string
	Title       string
	Author      string
	Cover       string
}

func (c *BuildCommand) Execute() {
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
