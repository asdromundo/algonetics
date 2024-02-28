/*
Copyright © 2024 Antonio S. Dromundo sebastiandromundo(at)outlook.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Executes a given genetic algorithm and profiles it.",
	Long: `The algonetics solve command executes a genetic algorithm chosen from a list of available options.
It is used to solve optimization problems by running the specified algorithm.

If no output file is provided, the command prints the results and performs profiling.
However, if an output file is specified, the command overwrites its contents with a text report summarizing the execution.
`,
	/*
		 	Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("solve called")
			},
	*/
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
