package main

import "fmt"

func _removeElem(pos int, arr []int) []int {
	var res []int
	res = append(arr[:pos], arr[pos+1:]...)
	return res
}

func main() {
	track := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(_removeElem(2, track))
}
