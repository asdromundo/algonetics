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
package genetic_algorithms

import (
	"fmt"
	"math/rand"
)

// ChessBoard representa un tablero de ajedrez para el problema de las N reinas.
type ChessBoard []uint

// Inicializa un tablero de ajedrez con N reinas colocadas de manera aleatoria.
func initializeBoard(n uint) ChessBoard {
	board := make(ChessBoard, n)
	// Creamos un slice para almacenar las filas que ya hemos utilizado
	usedRows := make([]bool, n)

	for i := uint(0); i < n; i++ {
		// Elegimos aleatoriamente una fila que no haya sido utilizada
		var row uint
		for {
			row = uint(rand.Intn(int(n)))
			if !usedRows[row] {
				usedRows[row] = true
				break
			}
		}
		board[i] = row
	}
	return board
}

func initializePopulation(n uint, populationSize uint) []ChessBoard {
	population := make([]ChessBoard, populationSize)
	for i := uint(0); i < populationSize; i++ {
		population[i] = initializeBoard(n) // n es el tamaño del tablero de ajedrez
	}
	return population
}

// Calcula la aptitud (fitness) de un tablero de ajedrez.
func calculateFitness(board ChessBoard) int {
	fitness := 0
	for i := 0; i < len(board); i++ {
		for j := i + 1; j < len(board); j++ {
			// Si dos reinas comparten la misma fila o columna, o están en la misma diagonal, hay una colisión
			if board[i] == board[j] || AbsInt(int(board[i])-int(board[j])) == j-i {
				fitness++
			}
		}
	}
	// La aptitud es inversamente proporcional al número de colisiones, es decir, menos colisiones significa mayor aptitud
	return len(board)*(len(board)-1)/2 - fitness
}

// Aplica el operador de cruce a la población de tableros de ajedrez.
func crossover(parents []ChessBoard, rate float64, algorithm string) []ChessBoard {
	offspring := make([]ChessBoard, 0)

	for i := 0; i < len(parents)-1; i += 2 {
		if rand.Float64() < rate {
			// Si elige cruzarse, se aplica el cruce según el algoritmo especificado
			switch algorithm {
			case "PMX":
				offspring = append(offspring, PMXCrossover(parents[i], parents[i+1]))
			// Agrega otros casos de algoritmos de cruce aquí según sea necesario
			default:
				fmt.Println("Algoritmo de cruce no válido")
				return nil
			}
		} else {
			// Si no se cruza, se mantienen los padres sin cambios
			offspring = append(offspring, parents[i], parents[i+1])
		}
	}

	return offspring
}

// Aplica elitismo conservando a los mejores individuos de una generación.
func applyElitism(population []ChessBoard, elites []ChessBoard) {
	// Copy the elites to the beginning of the population
	copy(population, elites)
}

// Verifica si se ha encontrado una solución al problema de las N reinas.
func hasSolution(population []ChessBoard) bool {
	for _, board := range population {
		if isSolution(board) {
			return true
		}
	}
	return false
}

// Función auxiliar para verificar si un tablero de ajedrez es una solución válida
func isSolution(board ChessBoard) bool {
	for i := 0; i < len(board); i++ {
		for j := i + 1; j < len(board); j++ {
			// Si dos reinas comparten la misma fila, columna o diagonal, no es una solución
			if board[i] == board[j] || AbsInt(int(board[i])-int(board[j])) == j-i {
				return false
			}
		}
	}
	return true
}

func NQueens(n uint, maxIterations uint, populationSize uint, crossoverRate float64, mutationRate float64, selectionMethod string, elitism bool, crossoverAlgorithm string) {
	// Lógica para resolver el problema de las N reinas
	fmt.Println("Running N Queens problem with the following parameters:")
	fmt.Printf("Number of Queens: %d\n", n)
	fmt.Printf("Max Iterations: %d\n", maxIterations)
	fmt.Printf("Population Size: %d\n", populationSize)
	fmt.Printf("Crossover Rate: %f\n", crossoverRate)
	fmt.Printf("Mutation Rate: %f\n", mutationRate)
	fmt.Printf("Selection Method: %s\n", selectionMethod)
	fmt.Printf("Elitism: %t\n", elitism)
	fmt.Printf("Crossover Algorithm: %s\n", crossoverAlgorithm)

	// Inicialización de la población
	population := initializePopulation(n, populationSize)

	// Iterar hasta alcanzar el número máximo de iteraciones
	for iteration := uint(0); iteration < maxIterations; iteration++ {
		// Iterar sobre cada individuo de la población
		for _, individual := range population {
			// Calcular la puntuación de aptitud (fitness) para el individuo
			calculateFitness(individual)
		}

		// Seleccionar individuos para la reproducción
		selectedParents := selectParents(population, selectionMethod)

		// Aplicar operador de cruce
		offspring := crossover(selectedParents, crossoverRate, crossoverAlgorithm)

		// Aplicar operador de mutación
		mutate(offspring, mutationRate)

		// Reemplazar la población actual con la descendencia generada
		population = offspring

		// Realizar elitismo si está habilitado
		if elitism {
			applyElitism(population, selectedParents)
		}

		// Verificar si se ha encontrado una solución
		if hasSolution(population) {
			fmt.Println("Solution found!")
			break
		}
	}

}
