package models

import "sort"

type Elves []*Elf

func (ees Elves) SortByMostCalories() {
	sort.SliceStable(ees, func(i, j int) bool {
		return ees[i].TotalCalories() > ees[j].TotalCalories()
	})
}

func (ees Elves) SortByMostItems() {
	sort.SliceStable(ees, func(i, j int) bool {
		return len(ees[i].Items) > len(ees[j].Items)
	})
}

func (ees Elves) SumOfCalories() int {
	result := 0
	for _, v := range ees {
		result += v.TotalCalories()
	}
	return result
}

func (ees Elves) SumOfItems() int {
	result := 0
	for _, v := range ees {
		result += len(v.Items)
	}
	return result
}

type Elf struct {
	Order int
	Items []int
}

func (e *Elf) AddItem(calorie int) {
	e.Items = append(e.Items, calorie)
}

func (e *Elf) TotalCalories() int {
	return sumItems(e.Items)
}
