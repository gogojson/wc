package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	count bool
	lines bool
	words bool
)

var rootCmd = &cobra.Command{
	Use:   "cc",
	Short: "This is a Word Count App",
	Run: func(cmd *cobra.Command, args []string) {
		var b []byte
		var fp string
		switch len(args) {
		case 0:
			ib, err := io.ReadAll(os.Stdin)
			if err != nil {
				panic("failed to read standard in")
			}
			b = ib
		default:
			var err error
			fp = args[0]
			b, err = os.ReadFile(fp)
			if err != nil {
				fmt.Printf("Cannot read file %s. Please check file path\n", fp)
				return
			}
		}

		switch {
		case count:
			getCount(b)
		case lines:
			getLines(b)
		case words:
			getWords(b)
		default:
			getLines(b)
			getWords(b)
			getCount(b)
		}

		switch fp {
		case "":
			fmt.Println()
		default:
			fmt.Println(fp)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&count, "count", "c", false, "Get byte count")
	rootCmd.PersistentFlags().BoolVarP(&lines, "lines", "l", false, "Get number of lines")
	rootCmd.PersistentFlags().BoolVarP(&words, "words", "w", false, "Get word count")

}

func getCount(b []byte) {
	fmt.Printf("%d ", len(b))
}

func getLines(b []byte) {
	var line int
	for _, b := range b {
		if string(b) == "\n" {
			line++
		}
	}
	fmt.Printf("%d ", line)
}

func getWords(b []byte) {
	fmt.Printf("%d ", len(bytes.Fields(b)))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
