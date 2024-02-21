package next_gen

import (
	"EGA_lab_10/structs"
	"math/rand"
	"sort"
)

func roulete(array []structs.Individual) int {

	index := 0

	var roulete, sum float32

	for i := range array {
		sum += array[i].Probability
	}
	shot := rand.Float32() * sum

	for ; roulete < shot && index < len(array)-1; index++ {
		roulete += array[index].Probability
	}

	return index

}

func Nextgen(distanceMatrix [][]structs.Way, previousgen []structs.Individual) []structs.Individual {

	prenextgen := PreNextgen(distanceMatrix, previousgen)
	var sum float32

	mutants := make([]structs.Individual, 0)
	mutants = append(mutants, prenextgen[0:len(previousgen)/2]...)

	sort.Slice(mutants, func(i, j int) bool {
		return mutants[i].Fitness < mutants[j].Fitness
	})

	parents := make([]structs.Individual, 0)
	parents = append(parents, prenextgen[0:len(previousgen)/2]...)

	for i := 1; i < len(mutants); i++ {
		kill := 0.4 + rand.Float32()*0.3

		if kill < mutants[i].Probability {
			mutants = append(mutants[:i], mutants[i+1:]...)
		}
	}

	for i := 0; i < len(mutants); i++ {
		sum += mutants[i].Fitness
	}

	for i := 0; i < len(mutants); i++ {
		mutants[i].Probability = 1 / mutants[i].Fitness
	}

	result := make([]structs.Individual, 0)

	for i := 0; i < len(previousgen)/2; i++ {
		result = append(result, mutants[roulete(mutants)])
	}

	result = append(result, parents...)

	return result
}
