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

import "math/rand"

func mutate(population []ChessBoard, rate float64) {
	for i := range population {
		if rand.Float64() < rate {
			// Si elige mutar, se aplica la mutación al tablero de ajedrez
			population[i] = mutation(population[i])
		}
	}
}

func mutation(board ChessBoard) ChessBoard {
	// Clonar el tablero original para realizar la mutación
	mutatedBoard := make(ChessBoard, len(board))
	copy(mutatedBoard, board)

	// Seleccionar dos posiciones aleatorias en el tablero para intercambiar las reinas
	position1 := rand.Intn(len(board))
	position2 := rand.Intn(len(board))

	// Intercambiar las reinas en las dos posiciones seleccionadas
	mutatedBoard[position1], mutatedBoard[position2] = mutatedBoard[position2], mutatedBoard[position1]

	return mutatedBoard
}
