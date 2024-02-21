package next_gen

import (
	"EGA_lab_10/structs"
	"math/rand"
)

func Mutation(individual structs.Individual) structs.Individual {
	var mutated = individual

	var position = rand.Intn(len(individual.Coding)-2) + 1
	mutated.Coding[position], mutated.Coding[len(individual.Coding)-1-position] = mutated.Coding[len(individual.Coding)-1-position], mutated.Coding[position]

	return mutated
}
