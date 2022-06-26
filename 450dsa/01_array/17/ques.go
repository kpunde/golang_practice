package _17

import (
	"reflect"
	"sort"
)

func nextPermutation(arr []int) []int {
	op := getAllPermutation(arr)
	refIndex := 0
	for index, item := range op {
		if reflect.DeepEqual(item, arr) {
			refIndex = index + 1
		}
	}

	return op[refIndex]
}

func recursivePerm(arr *[]int, res *[][]int, index int) {
	if index >= len(*arr) {
		var temp []int
		for _, item := range *arr {
			temp = append(temp, item)
		}
		*res = append(*res, temp)
		return
	}

	for i := index; i < len(*arr); i++ {
		(*arr)[index], (*arr)[i] = (*arr)[i], (*arr)[index]
		recursivePerm(arr, res, index+1)
		(*arr)[index], (*arr)[i] = (*arr)[i], (*arr)[index]
	}
}

func getAllPermutation(arr []int) [][]int {
	sort.Ints(arr)
	var result [][]int
	index := 0
	recursivePerm(&arr, &result, index)
	return result
}
