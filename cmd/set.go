package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/swarit-pandey/discrete-math/pkg/set"
)

var options set.Options

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set command is used to perform operations on sets",
	Long:  `Set can help you perform general operations on sets. For now only finite integral sets are supported.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate sets",
	Long:  `Generate sets of integers`,
	Run: func(cmd *cobra.Command, args []string) {
		if options.OutputFile == "" {
			options.OutputFile = createDefaultOutputFile()
		}

		set.NewSet().Generate(&options)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.AddCommand(generateCmd)

	// Flags specific to the generate subcommand
	generateCmd.Flags().IntVarP(&options.SetSize, "size", "", 10, "Size of each set")
	generateCmd.Flags().IntVarP(&options.NumberOfSets, "sets", "", 1, "Number of sets to generate")
	generateCmd.Flags().BoolVarP(&options.Randomize, "randomize", "r", false, "Randomize the elements in the set")
	generateCmd.Flags().IntVarP(&options.Range, "range", "R", 100, "Range of elements to generate")
	generateCmd.Flags().BoolVarP(&options.IgnoreDuplicate, "ignore-duplicate", "i", false, "Ignore duplicate elements in the set")
	generateCmd.Flags().StringVarP(&options.OutputFile, "output", "o", "", "Output file to write the sets to")
}

func createDefaultOutputFile() string {
	timestamp := time.Now().Format("20060102-150405")

	return fmt.Sprintf("sets-%s.json", timestamp)
}
