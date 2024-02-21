package startPopulation

import (
	"EGA_lab_10/structs"
	"sort"
)

func CreateStartpopulation(size int, distanceMatrix [][]structs.Way, numberOfcities int) []structs.Individual {

	startPopulation := make([]structs.Individual, size)

	for i := range startPopulation {
		startPopulation[i].Coding, startPopulation[i].Fitness = CreateIndividual(distanceMatrix, numberOfcities)
		startPopulation[i].Probability = 1 / startPopulation[i].Fitness
	}

	sort.Slice(startPopulation, func(i, j int) bool {
		return startPopulation[i].Coding[0] == startPopulation[i].Coding[len(startPopulation[i].Coding)-1]
	})

	return startPopulation
}
