package main

import "fmt"

func mean_median() {
	var data [100]int
	var n, sum int
	var avg float64

	fmt.Print("Masukkan Banyak Data: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan data : ")
		fmt.Scanln(&data[i])
		sum += data[i]
	}

	avg = float64(sum) / float64(n)
	fmt.Printf("Rata-rata dari data yang diinputkan adalah  %v \n", avg)

	if n%2 == 0 {
		median1 := (n / 2)
		median2 := (n / 2) - 1
		median := float64(data[median1]+data[median2]) / 2
		fmt.Printf("Median dari data yang diinputkan adalah %v \n", median)

	} else {
		median := (n - 1) / 2
		fmt.Printf("Median dari data yang diinputkan adalah %v \n", data[median])
	}
}
