package cmd

import (
	"github.com/janjitsu/gossaurus/internal/command"
	"github.com/spf13/cobra"
)

func buildCmd() *cobra.Command {
	buildCommand := command.BuildCommand{}
	build := &cobra.Command{
		Use:   "build",
		Short: "build dictionary file",
		Run: func(cmd *cobra.Command, args []string) {
			buildCommand.Execute()
		},
	}
	//required fields
	build.Flags().StringVar(&buildCommand.From, "from", "",
		"ISO two letter code of Origin Language. Ex. en")
	build.MarkFlagRequired("from")

	build.Flags().StringVar(&buildCommand.To, "to", "",
		"ISO two letter code of Target Language. Ex. es ")
	build.MarkFlagRequired("to")
	build.Flags().StringVar(&buildCommand.EntriesFile, "entries", "",
		"tsv file containing entries")
	build.MarkFlagRequired("entries")

	// optional fields
	build.Flags().StringVar(&buildCommand.Title, "title", "Converted Ebook",
		"Title of the Dictionary")
	build.Flags().StringVar(&buildCommand.Author, "author", "jan.art.br", "Name of author")
	build.Flags().StringVar(&buildCommand.Cover, "cover", "default-cover.jpg", "image for cover")

	return build
}

func init() {
	buildCmd := buildCmd()
	rootCmd.AddCommand(buildCmd)
}
