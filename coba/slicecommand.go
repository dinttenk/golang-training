package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputanCommand string
	var dataArray1 []string
	var dataArray2 []string
	dataArray3 := make([]int, 3, 3)

	fmt.Println("masukkan command string : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputanCommand = scanner.Text()
	dataArray1 = strings.Split(inputanCommand, ",")
	fmt.Println(len(dataArray1))
	for i := 0; i < len(dataArray1); i++ {
		dataArray2 = strings.Split(strings.Trim(dataArray1[i], " "), " ")
		switch dataArray2[0] {
		case "insert":
			posisi, _ := strconv.Atoi(dataArray2[1])
			value, _ := strconv.Atoi(dataArray2[2])
			dataArray3[posisi-1] = value
			fmt.Println(dataArray3)
		case "remove":
			posisi, _ := strconv.Atoi(dataArray2[1])
			dataArray3[posisi-1] = dataArray3[len(dataArray3)-1]
			dataArray3[len(dataArray3)-1] = 0
			dataArray3 = dataArray3[:len(dataArray3)-1]
			fmt.Println(dataArray3)
		case "append":
			value, _ := strconv.Atoi(dataArray2[1])
			fmt.Println(value)
			dataArray3 = append(dataArray3, value)
			fmt.Println(dataArray3)
		case "sort":
			sort.Ints(dataArray3)
			fmt.Println(dataArray3)
		case "pop":
			dataArray3[len(dataArray3)-1] = 0
			dataArray3 = dataArray3[:len(dataArray3)-1]
			fmt.Println(dataArray3)
		case "reverse":
			sort.Ints(dataArray3)
			sort.Slice(dataArray3, func(i, j int) bool {
				return dataArray3[i] > dataArray3[j]
			})
			fmt.Println(dataArray3)
		}
	}
}
