package main

import (
	"fmt"

	// "./aritmatika"
	"errors"
)

func main() {
	// hasilTambah := aritmatika.Tambah(1, 2)
	// fmt.Println(hasilTambah)
	name := "salman"
	_, err := validate(name)

	if err == nil {
		fmt.Println("hello, john")
	} else {
		fmt.Println("Error :", err)
	}
	fmt.Println("end")
}

func validate(name string) (bool, error) {
	if name != "" {
		return true, nil
	}
	return false, errors.New("nama tidak boleh ")
}
