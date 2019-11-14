package main


import "fmt"




// var name string = "john"

// var nameAddress *string = &name

// fmt.Println ("nama (value)", name)
// fmt.Println ("nama (address)", &name)

// fmt.Println ("nama (value)", *nameAddress)
// fmt.Println ("nama (address)", nameAddress)

// name = "doe"

// fmt.Println ("nama (value)", name)
// fmt.Println ("nama (address)", &name)

// fmt.Println ("nama (value)", *nameAddress)
// fmt.Println ("nama (address)", nameAddress)

// numberA := 25
// var numberB *int

// if numberB == nil {
// 	fmt.Println("numberB", numberB)
// 	numberB = &numberA
// 	fmt.Println ("numberB", *numberB)
// }
// var size = new (int)

// fmt.Printf ("Size (value): %d \n", *size)
// fmt.Printf ("Size (type): %T \n", size)
// fmt.Printf ("Size (value): %v \n", *&size)

// *size = 11
// fmt.Println ("New Size (value)", *size)

// hour := 15
// greeting (hour)


// }

// func greeting (hour int){

// 	if hour <12 {
// 		fmt.Println ("selamat jam lebih dari 12")
// 	}else {
// 		fmt.Println ("selamat")
// 	}


// p,l := 10,5
// luas,ket := calculate (p,l)

// fmt.Println (luas,ket)
// }

// // func calculate (panjang, lebar  int) int {
// // 	luas := panjang * lebar 
// // 	fmt.Println (luas,"1")
// // 	return luas
// // }

// func calculate (panjang, lebar  int) (luas int, ket string) {
// 	luas = panjang * lebar 
// 	ket = "keterangan"
	
// 	return 

// hasil := sum (1,2,3,4,5)
// fmt.Println (hasil)

// }

// func sum (numb ... int) int {
// 	var total int = 0
// 	banyakNumb := len(numb)
// 	for _, number := range numb {
// 		total += number
		
// 	}
// 	return total / banyakNumb


// func () {
// 	fmt.Println ("func pertama")
// } ()

// value := func() {
// 	fmt.Println("func kedua")
// }

// value ()
// func (sentence string){
// 	fmt.Println (sentence)
// } ("func ketiga")


func newCounter () func () int {
count :=0
return func() int {
	count += 1
	return count
}
}

func main (){
counter := newCounter()
fmt.Println (counter())
fmt.Println(counter())

}

