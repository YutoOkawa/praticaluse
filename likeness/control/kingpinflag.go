package control

import (
	"context"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	defaultLanguage = kingpin.Flag("default-language", "Default language").String()

	generateCmd = kingpin.Command("create-index", "Generate Index")
	inputFolder = generateCmd.Arg("INPUT", "Input Folder").Required().ExistingDir()

	searchCmd  = kingpin.Command("search", "Search")
	inputFile  = searchCmd.Flag("input", "Input index file").Short('i').File()
	searchWord = searchCmd.Arg("WORDS", "Search words").Strings()
)

func main() {
	ctx := context.Background()

	switch kingpin.Parse() {
	case generateCmd.FullCommand():
		err := generate(ctx)
		if err != nil {
			os.Exit(1)
		}
	case searchCmd.FullCommand():
		err := search(ctx)
		if err != nil {
			os.Exit(1)
		}
	}
}
