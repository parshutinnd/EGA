package next_gen

import (
	"EGA_lab_10/structs"
	"math/rand"
)

func PreNextgen(distanceMatrix [][]structs.Way, previousgen []structs.Individual) []structs.Individual {

	nextgen := make([]structs.Individual, len(previousgen))

	var sum float32

	parentsNumber := len(previousgen) / 2

	// супер особь
	var parent1Index, parent2Index int

	parent1Index = 0
	parent2Index = rand.Intn(parentsNumber - 1)

	if parent2Index == parent1Index {
		parent2Index++
	}

	nextgen[0], nextgen[1] = Crossover(previousgen[parent1Index], previousgen[parent2Index])

	for j := 0; j < len(nextgen[0].Coding)-1; j++ {
		nextgen[0].Fitness += distanceMatrix[nextgen[0].Coding[j]][nextgen[0].Coding[j+1]].Distanse
		nextgen[1].Fitness += distanceMatrix[nextgen[1].Coding[j]][nextgen[1].Coding[j+1]].Distanse
	}

	nextgen[0].Fitness = float32(int((nextgen[0].Fitness * 100))) / 100
	nextgen[1].Fitness = float32(int((nextgen[1].Fitness * 100))) / 100

	sum += nextgen[0].Fitness + nextgen[1].Fitness

	// все остальные
	for i := 2; i < parentsNumber; i += 2 {

		parent1Index = rand.Intn(parentsNumber - 1)
		parent2Index = rand.Intn(parentsNumber - 1)

		if parent2Index == parent1Index {
			parent2Index++
		}

		nextgen[i], nextgen[i+1] = Crossover(previousgen[parent1Index], previousgen[parent2Index])

		for j := 0; j < len(nextgen[i].Coding)-1; j++ {
			nextgen[i].Fitness += distanceMatrix[nextgen[i].Coding[j]][nextgen[i].Coding[j+1]].Distanse
			nextgen[i+1].Fitness += distanceMatrix[nextgen[i+1].Coding[j]][nextgen[i+1].Coding[j+1]].Distanse
		}

		nextgen[i].Fitness = float32(int((nextgen[i].Fitness * 100))) / 100
		nextgen[i+1].Fitness = float32(int((nextgen[i+1].Fitness * 100))) / 100

		sum += nextgen[i].Fitness + nextgen[i+1].Fitness

	}

	for i := 0; i < parentsNumber; i += 2 {
		mutationPropability := rand.Intn(10)
		if mutationPropability < 2 {
			nextgen[i] = Mutation(nextgen[i])

			for j := 0; j < len(nextgen[i].Coding)-1; j++ {
				nextgen[i].Fitness += distanceMatrix[nextgen[i].Coding[j]][nextgen[i].Coding[j+1]].Distanse
			}
			nextgen[i].Fitness = float32(int((nextgen[i].Fitness * 100))) / 100
		}
	}

	for i := parentsNumber; i < len(previousgen); i++ {
		nextgen[i] = previousgen[len(previousgen)-i]
	}

	for i := range previousgen {
		nextgen[i].Probability = 1 / nextgen[i].Fitness
	}
	return nextgen
}
