package startPopulation

import (
	"EGA_lab_10/structs"
	"math/rand"
)

func NextCity(candidates []structs.Way, cities []bool) (int, float32) {

	var roulette, sum, shot float32

	for i := range candidates {
		if !candidates[i].Chosed && !cities[i] {
			sum += candidates[i].Probability
		}
	}

	shot = rand.Float32() * sum

	i := 0

	for {

		if !candidates[i].Chosed && !cities[i] {
			roulette += candidates[i].Probability

			if roulette >= shot {
				break
			}

			i++
		} else {
			i++
		}
	}

	return i, candidates[i].Distanse
}

func CreateIndividual(inputMatrix [][]structs.Way, numberOfcities int) ([]int, float32) {

	cities := make([]bool, numberOfcities)

	distanceMatrix := make([][]structs.Way, numberOfcities)

	for i := range distanceMatrix {

		distanceMatrix[i] = make([]structs.Way, numberOfcities)
	}

	for i := range distanceMatrix {
		for j := 0; j < numberOfcities; j++ {
			distanceMatrix[i][j] = inputMatrix[i][j]
		}
	}

	i := 1

	resultObhod := make([]int, 0)

	currentCityinex := rand.Intn(numberOfcities)
	resultObhod = append(resultObhod, currentCityinex)
	cities[currentCityinex] = true

	var distanceFromprevious, result float32

	for len(resultObhod) < numberOfcities {

		i++

		nextCitycandidate1, distance1 := NextCity(distanceMatrix[resultObhod[len(resultObhod)-1]], cities)
		nextCitycandidate2, distance2 := NextCity(distanceMatrix[resultObhod[0]], cities)

		var nextCity int

		if distance1 <= distance2 {
			nextCity = nextCitycandidate1
			resultObhod = append(resultObhod, nextCity)
			distanceFromprevious = distance1
		} else {
			nextCity = nextCitycandidate2
			resultObhod = append([]int{nextCity}, resultObhod...)
			distanceFromprevious = distance2
		}

		if len(resultObhod) == len(cities) {
			nextCity = resultObhod[0]
			result += distanceMatrix[resultObhod[len(resultObhod)-1]][resultObhod[0]].Distanse
			resultObhod = append(resultObhod, nextCity)

		}

		cities[nextCity] = true

		distanceMatrix[currentCityinex][nextCity].Chosed = true
		distanceMatrix[nextCity][currentCityinex].Chosed = true

		result += distanceFromprevious

		currentCityinex = nextCity
	}

	return resultObhod, float32(int((result * 100))) / 100

}
