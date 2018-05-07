package main

import (
	"testing"
	"fmt"
)

func TestMapBasic(t *testing.T) {
	dict := make(map[string]int)
	dict = map[string]int{"hello": 1, "world": 2}
	fmt.Println(dict)

	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict2)

	dict2["Blue"] = "#dea123"
	fmt.Println(dict2)

	value, exists := dict2["Blue"]
	if (exists) {
		fmt.Println(value)
	}
	println(dict2)
	removeKey(dict2, "Red")
	fmt.Println("\n")
	println(dict2)

}

func println(dict map[string]string) {
	for key, val := range dict {
		fmt.Printf("Key:%s,Value:%s\t", key, val)
	}
}
func removeKey(dict map[string]string, key string) {
	delete(dict, key)
}
