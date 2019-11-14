package main

import "fmt"

func main() {

	key1 := []string{"firstname", "lastname", "age"}

	key2 := []string{"rochmadn", "nurdin", "257"}

	csvmap := make(map[string]string)
	for i := 0; i < len(key1); i++ {
		csvmap[key1[i]] = key2[i]
	}
	fmt.Println(csvmap)
}
