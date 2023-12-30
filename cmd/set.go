package cmd

import (
	"github.com/spf13/cobra"
	"github.com/swarit-pandey/discrete-math/pkg/set"
)

var options set.Options

func newSetCmd() *cobra.Command {
	setCmd := cobra.Command{
		Use:   "set",
		Short: "perform ops on integral sets",
		Long:  "Set can help you perform basic set operations",
	}

	setCmd.AddCommand(newGenerateCmd())
	setCmd.AddCommand(newUnionCmd())
	setCmd.AddCommand(newIntersectCmd())
	setCmd.AddCommand(newPowersetCmd())

	return &setCmd
}

func newGenerateCmd() *cobra.Command {
	generateCmd := cobra.Command{
		Use:   "generate",
		Short: "Generate integral sets",
		Long:  "Generate sets",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := set.NewSet().Generate(&options)
			if err != nil {
				return err
			}

			return nil
		},
	}

	generateCmd.Flags().IntVarP(&options.SetSize, "size", "", 10, "Size of each set")
	generateCmd.Flags().IntVarP(&options.NumberOfSets, "sets", "", 1, "Number of sets to generate")
	generateCmd.Flags().BoolVarP(&options.Randomize, "randomize", "r", false, "Randomize the elements in the set")
	generateCmd.Flags().IntVarP(&options.Range, "range", "R", 100, "Range of elements to generate")
	generateCmd.Flags().BoolVarP(&options.IgnoreDuplicate, "ignore-duplicate", "", false, "Ignore duplicate elements in the set")
	generateCmd.Flags().StringVarP(&options.OutputFile, "output", "o", "", "Output file to write the sets to")
	generateCmd.Flags().StringVarP(&options.InputFile, "input", "i", "", "Input json file path")

	return &generateCmd
}

func newUnionCmd() *cobra.Command {
	unionCmd := cobra.Command{
		Use:   "union",
		Short: "get union of sets",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := set.NewSet().Union(&options)
			if err != nil {
				return err
			}

			return nil
		},
	}

	unionCmd.Flags().StringVarP(&options.InputFile, "input", "", "", "Input json file path")
	unionCmd.Flags().StringVarP(&options.OutputFile, "output", "", "", "Output resultant file")

	return &unionCmd
}

func newIntersectCmd() *cobra.Command {
	intersectCmd := cobra.Command{
		Use:   "intersect",
		Short: "Get intersection of sets",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := set.NewSet().Intersect(&options)
			if err != nil {
				return err
			}

			return nil
		},
	}

	intersectCmd.Flags().StringVarP(&options.InputFile, "input", "", "", "Input json file path")
	intersectCmd.Flags().StringVarP(&options.OutputFile, "output", "", "", "Output json file path")

	return &intersectCmd
}

func newPowersetCmd() *cobra.Command {
	powersetCmd := cobra.Command{
		Use:   "powerset",
		Short: "generate powerset of sets",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := set.NewSet().Powerset(&options)
			if err != nil {
				return err
			}

			return nil
		},
	}

	powersetCmd.Flags().StringVarP(&options.InputFile, "input", "", "", "Input json file path")
	powersetCmd.Flags().StringVarP(&options.OutputFile, "output", "", "", "Output json file path")

	return &powersetCmd
}

func init() {
	rootCmd.AddCommand(newSetCmd())
}
