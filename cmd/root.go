package cmd

import (
	"fmt"
	"os"
	"strings"

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
		if len(args) < 1 {
			panic("Needs to have at lest one argument file path")
		}
		fp := args[0]

		if count {
			getCount(fp)
		}
		if lines {
			getLines(fp)
		}
		if words {
			getWords(fp)
		}

		if !count && !lines && !words {
			getCount(fp)
			getLines(fp)
			getWords(fp)
		}

		fmt.Println(fp)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&count, "count", "c", false, "Get byte count")
	rootCmd.PersistentFlags().BoolVarP(&lines, "lines", "l", false, "Get number of lines")
	rootCmd.PersistentFlags().BoolVarP(&words, "words", "w", false, "Get word count")

}

func getCount(fp string) {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Cannot read file %s. Please check file path\n", fp)
		return
	}
	fmt.Printf("%d ", len(b))
}

func getLines(fp string) {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Cannot read file %s. Please check file path\n", fp)
		return
	}

	var line int

	for _, b := range b {
		if string(b) == "\n" {
			line++
		}
	}
	fmt.Printf("%d ", line+1)
}

func getWords(fp string) {
	b, err := os.ReadFile(fp)
	if err != nil {
		fmt.Printf("Cannot read file %s. Please check file path\n", fp)
		return
	}

	var word int

	for _, b := range b {
		//TODO: Do not count when the first character is new line or space
		//TODO: Do not count when there are multiple new line or space in one row

		// if word == 0 && strings.Contains("\n., ", string(b)) {
		// 	fmt.Println("first is bla")
		// 	continue
		// }
		if strings.Contains("\n., ", string(b)) {
			word++
		}
	}
	fmt.Printf("%d ", word+1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
