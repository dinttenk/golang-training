package main

import (
	"fmt"
	"strings"
)

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func main() {
// 	var Person0 Person
// 	Person0.FirstName = "muchson"
// 	Person0.LastName = "Ibi"
// 	Person0.Age = 27

// 	fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)

// 	var person1 = Person{"Rizky", "Kurniawan", 26}
// 	fmt.Println(person1)

// 	var Person2 = Person{
// 		FirstName: "iswanul",
// 		LastName:  "umam",
// 		Age:       25,
// 	}
// 	fmt.Println(Person2)

// 	Person3 := Person{"Pranadya", "bagus", 23}
// 	fmt.Println(Person3)

// 	Person4 := new(Person)
// 	Person4.FirstName = "Muhammad"
// 	Person4.LastName = "ismail"
// 	Person4.Age = 30
// 	fmt.Println(*Person4)

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// type Person2 struct {
// 	FirstName string
// 	LastName  string
// 	Tangan    []Tangan
// }

// type Tangan struct {
// 	Warna   string
// 	Panjang int
// }

// func main() {
// 	var person5 = Person2{
// 		FirstName: "Ila",
// 		LastName:  "umam",
// 		Tangan: []Tangan{
// 			Tangan{
// 				Warna:   "Kecoklatan",
// 				Panjang: 15,
// 			},
// 			Tangan{
// 				Warna:   "Kuning langsat",
// 				Panjang: 40,
// 			},
// 		},
// 	}
// 	fmt.Println(person5)

// type Person struct {
// 	name string
// 	age  int
// }

// func (P Person) GetName() string {
// 	return P.name + "amazing!"
// }
// func (P *Person) IncreaseAge() {
// 	P.age = P.age + 1
// }
// func main() {
// 	PersonA := Person{"john", 50}
// 	fmt.Printf("%v\n", PersonA)
// 	fmt.Println(PersonA.GetName())

// 	PersonA.IncreaseAge()
// 	fmt.Println(PersonA.age)

// type Rect struct {
// 	width  float64
// 	height float64
// 	area   float64
// }
// type circle struct {
// 	radius float64
// }

// func (r Rect) Area() float64 {
// 	return r.width * r.height

// }
// func (r *Rect) CalculateArea() float64 {
// 	r.area = r.width * r.height

// }
// func (c circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func main() {
// 	fmt.Println("")
// 	Rect := Rect{
// 		width:  5.0,
// 		height: 4.0}
// 	Cir := circle{5.0}
// 	fmt.Printf("Area rectangle = %0.1f\n", Rect.Area())
// 	fmt.Printf(" Area circle = %f\n", Cir.Area())

// 	fmt.Printf("Rectangle before call pointer method = %v\n", Rect)
// 	Rect.CalculateArea()
// 	fmt.Printf("Rectangle acter call pointer method = %v\n", Rect)
// 	fmt.Println("")

// }

// type square struct {
// 	side int
// }
// type calculate interface {
// 	large() int
// 	doubleSide() int
// 	nSide(n int) int
// }

// func main() {
// 	var dimResult calculate
// 	dimResult = square{10}
// 	fmt.Println("Large square", dimResult.large())
// 	fmt.Println("Large square", dimResult.doubleSide())
// }

// func (s square) large() int {
// 	return s.side * s.side
// }
// func (s square) doubleSide() int {
// 	return s.side * 2
// }
// func (s square) nSide(n int) int {
// 	return s.side * n
// }

func main() {
	fmt.Println()
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multipled by 10 is :", number)

	secret = []string{"apple", "mango", "banana"}
	var gruits = strings.Join(secret.([]string), ",")
	fmt.Println(gruits, "is may favorite fruits")

	secret = 45
	explain(secret)
	fmt.Println()
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored as string", i.(string))
	case int:
		fmt.Println("i stored int", i.(int))
	case bool:
		fmt.Println("i stored boolean", i.(bool))
	default:
		fmt.Println("i stored as something else", i)

	}
}
