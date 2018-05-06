package gotest

import "testing"

func TestDeclare(test *testing.T) {
	var array = [...]int{10, 20, 30, 40, 50}
	array = [5]int{1: 20, 2: 50}
	println(array[1])
	var array2 = [5]int{10, 20, 30, 40, 50}
	println(array2[1])
}