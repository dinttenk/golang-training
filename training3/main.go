package main

import "fmt"

func main() {

	//faktor bilangan

	var tambah int
	var n int

	fmt.Println ("Masukan angka :")
	fmt.Scanf("%v", &n)
	for i:=0; i<=n; i++{
		tambah++
		if n% tambah == 0 {
				fmt.Println(tambah)
	}
	}
	}

	// bilangan prima
	var angka int

	var prima bool = true
	var x int

	fmt.Println("Masukan angka :")
	fmt.Scanln(&angka)

	for pembagi := 2; pembagi <= angka/2; pembagi++ {
		x = angka % pembagi

		if x == 0 {
			prima = false
			break
		}
	}
	if prima && angka > 0 && angka != 0 {
		fmt.Println("adalah bilangan prima")
	} else {
		fmt.Println("bukan bilangan prima")

	}
}

//palindrome

	var poli string
	var input string

	fmt.Println("Masukan angka :")
	fmt.Scanln(&input)

	for i := len(input); i > 0; i-- {
		poli += string(input[i-1])

	}
	if poli == input {
		fmt.Println("pali cuy")
	} else {
		fmt.Println("bukan pali cuy")
	}

	var x, y, z int
	var n int

	fmt.Println("Masukan angka :")
	fmt.Scanf("%v", &n)

	for x = 1; x <= n; x++ {

		for y = n; y >= x; y-- {
			fmt.Print(" ")
		}
		for z = 1; z <= x; z++ {
			fmt.Print("* ")
		}
		// for a=1; a<=x-1; a++ {
		// 	fmt.Print ("*")

		// }
		fmt.Println(" ")

	}
}
