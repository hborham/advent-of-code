package models

func sumItems(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func SumElvesTotalCaloriesOld(array Elves) int {
	result := 0
	for _, v := range array {
		result += v.TotalCalories()
	}
	return result
}

func SumElvesTotalItemsOld(array Elves) int {
	result := 0
	for _, v := range array {
		result += len(v.Items)
	}
	return result
}
