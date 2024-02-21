package main

import (
	"EGA_lab_10/next_gen"
	"EGA_lab_10/startPopulation"
	"EGA_lab_10/structs"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	numberOfcities := 15

	distanceMatrix := make([][]structs.Way, numberOfcities)

	for i := range distanceMatrix {

		distanceMatrix[i] = make([]structs.Way, numberOfcities)
	}

	var lines []string
	file, _ := os.Open("text.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := range distanceMatrix {
		arr := strings.Split(lines[i], " ")
		for j := i; j < numberOfcities; j++ {
			if i != j {
				element, _ := strconv.ParseFloat(arr[j], 32)
				distanceMatrix[i][j].Distanse = float32(element)
				distanceMatrix[i][j].Probability = 1.0 / float32(distanceMatrix[i][j].Distanse)
				distanceMatrix[j][i] = distanceMatrix[i][j]
			} else {
				distanceMatrix[i][j].Chosed = true
			}
		}
	}

	for i := range distanceMatrix {
		for j := range distanceMatrix[i] {
			fmt.Print(distanceMatrix[i][j].Distanse, " ")
		}
		fmt.Println()
	}

	fmt.Println()

	sizePopulation := 30

	fmt.Println("Start population:  ")
	fmt.Println()

	startPopulation := startPopulation.CreateStartpopulation(sizePopulation, distanceMatrix, numberOfcities)

	sort.Slice(startPopulation, func(i, j int) bool {
		return startPopulation[i].Fitness < startPopulation[j].Fitness
	})

	for i := range startPopulation {
		fmt.Println(startPopulation[i].Coding, " ", startPopulation[i].Fitness)
	}

	fmt.Println()

	var bestIndiviual structs.Individual
	bestIndiviual.Fitness = 100

	for i := 0; i < 20; i++ {

		fmt.Println("Nomer pokoleniya ", i+1)
		fmt.Println()

		nextgen := next_gen.Nextgen(distanceMatrix, startPopulation)

		sort.Slice(nextgen, func(i, j int) bool {
			return nextgen[i].Fitness < nextgen[j].Fitness
		})

		for i := range nextgen {
			fmt.Println(nextgen[i].Coding, " ", nextgen[i].Fitness)
		}

		fmt.Println()
		fmt.Println("best indiviual from popultion: ", nextgen[0].Coding, " ", nextgen[0].Fitness)

		if nextgen[0].Fitness < bestIndiviual.Fitness {
			bestIndiviual = nextgen[0]
		}

		fmt.Println()
	}

	fmt.Println("best iniviual: ", bestIndiviual.Coding, " ", bestIndiviual.Fitness)

}
