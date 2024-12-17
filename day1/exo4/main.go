package main

import "fmt"

func main() {
	mapTest1 := map[string]int{
		"1": 1,
		"2": 1,
		"3": 1,
	}

	mapTest2 := map[string]int{
		"1": 1,
		"4": 1,
		"5": 1,
	}

	result := fusionMaps(mapTest1, mapTest2)

	fmt.Printf("%v", result)
}

func fusionMaps(map1, map2 map[string]int) map[string]int {
	result := map1

	for key, value := range map2 {
		_, exist := result[key]

		if exist {
			result[key] += value
		} else {
			result[key] = value
		}
	}

	return result
}
