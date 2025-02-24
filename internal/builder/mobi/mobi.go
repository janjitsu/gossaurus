package mobi

import (
	"fmt"
	"os"

	"github.com/janjitsu/gossaurus/internal/entities"
)

func BuildPage(page entities.Page, target string) {
	baseName := page.Id
	extension := "html"
	// Create a filename by combining baseName and extension
	filename := fmt.Sprintf("out/%s/%s.%s", target, baseName, extension)
	fmt.Println(filename)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = RenderTemplate(f, "page", page)
	if err != nil {
		panic(err)
	}
}

func BuildDictionaryContent(dictionary entities.Dictionary) {
	filename := fmt.Sprintf("out/%s/content.opf", dictionary.Id)
	fmt.Println(filename)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	//@TODO salvar o html para o arquivo
	err = RenderTemplate(f, "content", dictionary)
	if err != nil {
		panic(err)
	}
}
