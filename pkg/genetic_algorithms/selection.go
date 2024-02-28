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

// Selecciona padres para la reproducción utilizando el método especificado.
func selectParents(population []ChessBoard, method string) []ChessBoard {
	switch method {
	case "roulette":
		return rouletteSelection(population)
	case "tournament":
		return tournamentSelection(population)
	default:
		fmt.Println("Método de selección no válido")
		return nil
	}
}

// Selección de padres por ruleta
func rouletteSelection(population []ChessBoard) []ChessBoard {
	// Calcular la suma total de la aptitud de la población
	totalFitness := 0
	for _, board := range population {
		totalFitness += calculateFitness(board)
	}

	// Seleccionar padres aleatorios basados en la probabilidad proporcional a su aptitud
	selectedParents := make([]ChessBoard, 0)
	for len(selectedParents) < len(population) {
		// Seleccionar un padre aleatorio
		selectedParentIndex := selectParentIndex(population, totalFitness)
		selectedParent := population[selectedParentIndex]
		selectedParents = append(selectedParents, selectedParent)
	}

	return selectedParents
}

// Función auxiliar para seleccionar un índice de padre basado en la probabilidad proporcional a su aptitud
func selectParentIndex(population []ChessBoard, totalFitness int) int {
	// Generar un valor aleatorio entre 0 y la suma total de la aptitud
	randFitness := rand.Intn(totalFitness)

	// Elegir el índice de un padre basado en la probabilidad proporcional a su aptitud
	cumulativeFitness := 0
	for i, board := range population {
		cumulativeFitness += calculateFitness(board)
		if cumulativeFitness > randFitness {
			return i
		}
	}

	// En caso de que algo salga mal, devolvemos el último índice
	return len(population) - 1
}

// Selección de padres por torneo
func tournamentSelection(population []ChessBoard) []ChessBoard {
	// Seleccionar dos individuos aleatorios para cada torneo y elegir al mejor
	var selectedParents []ChessBoard
	for i := 0; i < len(population); i += 2 {
		parent1 := population[rand.Intn(len(population))]
		parent2 := population[rand.Intn(len(population))]
		if calculateFitness(parent1) > calculateFitness(parent2) {
			selectedParents = append(selectedParents, parent1)
		} else {
			selectedParents = append(selectedParents, parent2)
		}
	}

	return selectedParents
}
