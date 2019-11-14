package main


import (
	"fmt"
	
)

func main(){
	
	//     var 
	// 	var num[100] int
	// 	var temp,sum,avg int
	// 	fmt.Print("Enter number of elements: ")
	// 	fmt.Scanln(&temp)
	// 	for i := 0; i < temp; i++ {
	// 		fmt.Print("Enter the number : ")
	// 		fmt.Scanln(&num[i])     
	// 		sum += num[i]
	// 	}
		 
	// 	avg = sum/temp
	// 	fmt.Printf("The Average of entered %d number(s) is %d",temp,avg)

	// }
	
		


		//bool default value false
		// slice := append ([]string {"didin","ayu","sinta","didin", "ayu"},) 
		// slice_2 := append ([]string {"didin","ayu","sinta","didin", "ayu"}) 
		
var temp int




// 	slice_2 := append ([]string {}) 
	
	
	fmt. Print ("enter jumlah data ")
	fmt.Scanln (&temp)
	var slice  = new ([temp]string )
	var slice_2  = new ( [temp]string)

	for i:=0; i<temp; i++{
		fmt.Print("masukan kata")
		fmt.Scanln(&slice[i])
	}
	for i:=0; i<temp; i++{
		fmt.Print("masukan kata")
		fmt.Scanln(&slice_2[i])

	
	}
	slice_3 := append(slice,slice_2...)
	keys := make (map [string]bool)
	list := [] string {}
	

}

	// for _, n := range slice_3 { // loop cek slice
	// 	if _, value := keys [n]; !value { // seleksi jika n (slice) tidak sama sama value 
	// 		keys [n]= true  // range size true
	// 		list = append (list,n) 
	// 	}
	// }
	// fmt.Println (list)


// a := []int{ 1, 3, 5 ,4,5}
// n := len(a)
// var avg int
// sum:= 0
// for i:=0; i<n; i++{
// 	sum +=a[i]
	
	
// }
// fmt.Println(avg)







	
	
