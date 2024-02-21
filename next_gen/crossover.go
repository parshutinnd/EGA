package next_gen

import (
	"EGA_lab_10/structs"
	"math/rand"
)

func deleteValues(slice []int, values []int, startValue int) []int {

	for i := range values {
		for j := startValue; j < len(slice); j++ {
			if slice[j] == values[i] {
				slice = append(slice[:j], slice[j+1:]...)
			}
		}
	}

	return slice
}

func Crossover(individual1, individual2 structs.Individual) (structs.Individual, structs.Individual) {

	point := rand.Intn(len(individual1.Coding)-3) + 1

	var result1, result2 structs.Individual

	result1.Coding = append(result1.Coding, individual2.Coding[0])
	result1.Coding = append(result1.Coding, individual2.Coding[point:point+2]...)

	result1.Coding = append(result1.Coding, individual1.Coding[1:len(individual2.Coding)-1]...)

	values1 := make([]int, 0)
	values1 = append(values1, result1.Coding[0:3]...)
	result1.Coding = deleteValues(result1.Coding, values1, 3)

	if len(result1.Coding) == len(individual1.Coding)-2 {
		result1.Coding = append(result1.Coding, individual1.Coding[0])
	}
	result1.Coding = append(result1.Coding, individual2.Coding[0])
	//result1 done here

	result2.Coding = append(result2.Coding, individual1.Coding[0])
	result2.Coding = append(result2.Coding, individual1.Coding[point:point+2]...)

	result2.Coding = append(result2.Coding, individual2.Coding[1:len(individual1.Coding)-1]...)

	values2 := make([]int, 0)
	values2 = append(values2, result2.Coding[0:3]...)
	result2.Coding = deleteValues(result2.Coding, values2, 3)

	if len(result2.Coding) == len(individual2.Coding)-2 {
		result2.Coding = append(result2.Coding, individual2.Coding[0])
	}
	result2.Coding = append(result2.Coding, individual1.Coding[0])
	//result2 done here

	return result1, result2

}
