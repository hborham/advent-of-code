package main

import (
	"advent-of-code/models"
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

type (
	Elves = models.Elves
	Elf   = models.Elf
)

func main() {
	filename := "input.txt"

	txtlines := utils.ReadFile(filename)

	elves := getElves(txtlines)

	elves.SortByMostCalories()

	top := 3

	topElf := elves[0]
	fmt.Println("Top by calories")
	fmt.Printf("Elf: %d is carrying %d items with %d calories\n", topElf.Order, len(topElf.Items), topElf.TotalCalories())
	topElves := elves[0:top]
	fmt.Printf("Top 3 elves are carrying %d items totaling %d calories\n", topElves.SumOfItems(), topElves.SumOfCalories())

	elves.SortByMostItems()

	topElf = elves[0]
	fmt.Println("Top by items")
	fmt.Printf("Elf: %d is carrying %d items with %d calories\n", topElf.Order, len(topElf.Items), topElf.TotalCalories())
	topElves = elves[0:top]
	fmt.Printf("Top 3 elves are carrying %d items totaling %d calories\n", topElves.SumOfItems(), topElves.SumOfCalories())
}

func getElves(items []string) Elves {
	var order = 1
	var elves Elves
	var elf = new(Elf)
	elf.Order = order
	for i := 0; i < len(items); i++ {
		n, err := strconv.Atoi(items[i])
		if err != nil {
			order++
			elves = append(elves, elf)
			elf = new(Elf)
			elf.Order = order
			continue
		}
		elf.AddItem(n)

	}
	return elves
}
