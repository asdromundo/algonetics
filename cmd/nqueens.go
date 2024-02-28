/*
Copyright © 2024 Antonio S. Dromundo sebastiandromundo(at)outlook.com
*/
package cmd

import (
	"asdrome.com/algonetics/pkg/genetic_algorithms"
	"github.com/spf13/cobra"
)

var numberQueens uint
var crossoverAlgorithm string

// nqueensCmd represents the nqueens command
var nqueensCmd = &cobra.Command{
	Use:   "nqueens",
	Short: "The N Queens Problem Algorithm",
	Long: `The n-queens problem is to place n queens (where n > 0) on an n-by-n chessboard so that no queen is threatened by another one.

According to the rules of chess, this is equivalent to the requirement that no two queens be on the same row or the same column or on a common diagonal.`,
	Run: func(cmd *cobra.Command, args []string) {
		runGeneticAlgorithm()
	},
}

func init() {
	solveCmd.AddCommand(nqueensCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nqueensCmd.PersistentFlags().String("foo", "", "A help for foo")

	nqueensCmd.PersistentFlags().UintVarP(&numberQueens, "number-queens", "n", 8, "Positive integer representing the number of queens.")
	nqueensCmd.PersistentFlags().StringVarP(&crossoverAlgorithm, "crossover-algorithm", "a", "pmx", "Crossover algorithm to use (pmx, ipmx, opmx)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nqueensCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func runGeneticAlgorithm() {

	// Leer configuraciones de las banderas globales
	maxIterations, _ := rootCmd.PersistentFlags().GetUint("max-iterations")
	populationSize, _ := rootCmd.PersistentFlags().GetUint("population-size")
	crossoverRate, _ := rootCmd.PersistentFlags().GetFloat64("crossover-rate")
	mutationRate, _ := rootCmd.PersistentFlags().GetFloat64("mutation-rate")
	selectionMethod, _ := rootCmd.PersistentFlags().GetString("selection-method")
	elitism, _ := rootCmd.PersistentFlags().GetBool("elitism")

	// Ejecutar el algoritmo genético con las configuraciones proporcionadas
	// Aquí puedes llamar a la función que implementa el algoritmo de las reinas n
	// Pasando las configuraciones como parámetros
	genetic_algorithms.NQueens(numberQueens, maxIterations, populationSize, crossoverRate, mutationRate, selectionMethod, elitism, crossoverAlgorithm)
}
