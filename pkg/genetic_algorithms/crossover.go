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
	"math/rand"
)

// PMX aplica el operador de cruce PMX (Partially Mapped Crossover)
func PMX(parent1, parent2 ChessBoard) ChessBoard {
	size := len(parent1)
	// Crea un descendiente con la misma longitud que los padres
	descendant := make(ChessBoard, size)

	// Selecciona dos puntos de cruce aleatorios
	crossoverPoint1 := rand.Intn(size)
	crossoverPoint2 := rand.Intn(size)
	for crossoverPoint2 == crossoverPoint1 {
		crossoverPoint2 = rand.Intn(size)
	}
	if crossoverPoint2 < crossoverPoint1 {
		crossoverPoint1, crossoverPoint2 = crossoverPoint2, crossoverPoint1
	}

	// Copia la sección entre los puntos de cruce del primer padre al descendiente
	copy(descendant[crossoverPoint1:crossoverPoint2], parent1[crossoverPoint1:crossoverPoint2])

	// Realiza el mapeo parcial entre el primer padre y el segundo padre
	for i := crossoverPoint1; i < crossoverPoint2; i++ {
		// Si el valor actual del descendiente proviene del primer padre, busca su equivalente en el segundo padre
		if descendant[i] == parent1[i] {
			val := parent2[i]
			for val != parent1[i] {
				// Encuentra la posición de este valor en el primer padre
				pos := -1
				for j := 0; j < size; j++ {
					if parent1[j] == val {
						pos = j
						break
					}
				}
				// Si el valor no es igual al que se está buscando, continua buscando
				if pos == -1 {
					descendant[i] = val
					val = parent2[pos]
				}
			}
		}
	}

	// Completa el descendiente con los valores restantes del segundo padre
	for i := 0; i < len(descendant); i++ {
		if descendant[i] == 0 {
			descendant[i] = parent2[i]
		}
	}

	return descendant
}
