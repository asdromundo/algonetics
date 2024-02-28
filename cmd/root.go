/*
Copyright © 2024 Antonio S. Dromundo sebastiandromundo(at)outlook.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "algonetics",
	Short: "CLI for testing, profiling and reporting genetic algorithms implementations in Go(lang).",
	Long: `
algonetics is a command-line interface (CLI) tool designed for testing, profiling, and reporting genetic algorithms implementations in the Go programming language. This tool empowers researchers, developers, and enthusiasts to experiment with various genetic algorithm configurations, including selection strategies, crossover techniques, and mutation rates, providing insights into their performance and effectiveness across different problem domains.

By leveraging the power of Go(lang) concurrency and the flexibility of the Cobra library, algonetics offers a seamless and efficient environment for analyzing genetic algorithms, enabling users to tackle optimization problems with ease.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.algonetics.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().UintP("max-iterations", "i", 100, "Maximum number of iterations")
	rootCmd.PersistentFlags().UintP("population-size", "p", 10, "Population size")
	rootCmd.PersistentFlags().Float64P("crossover-rate", "c", 0.8, "Crossover rate")
	rootCmd.PersistentFlags().Float64P("mutation-rate", "m", 0.01, "Mutation rate")
	rootCmd.PersistentFlags().StringP("selection-method", "s", "roulette", "Selection method")
	rootCmd.PersistentFlags().BoolP("elitism", "e", false, "Enable elitism")

	// Bind flags to viper
	viper.BindPFlag("max-iterations", rootCmd.PersistentFlags().Lookup("max-iterations"))
	viper.BindPFlag("population-size", rootCmd.PersistentFlags().Lookup("population-size"))
	viper.BindPFlag("crossover-rate", rootCmd.PersistentFlags().Lookup("crossover-rate"))
	viper.BindPFlag("mutation-rate", rootCmd.PersistentFlags().Lookup("mutation-rate"))
	viper.BindPFlag("selection-method", rootCmd.PersistentFlags().Lookup("selection-method"))
	viper.BindPFlag("elitism", rootCmd.PersistentFlags().Lookup("elitism"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".algonetics" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".algonetics")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
