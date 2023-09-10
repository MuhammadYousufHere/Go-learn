// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// var greet string = "Hello Yousuf"
// var age int = 25

// const name string = "Muhammad Yousuf"
// const pi float32 = 3.14
// const goLearner bool = true

// // Arrays

// var arr  [3]int
// // arr[1] = 4

// // array litrals

// var numbers [5]int = [5]int{2,3,4,5,6} // arrays with five items

// // Slices

// var vowels = [...]string{"a","i","e","o","u"}
// // Slices using make()
// //  type - length of slice - cap of slice
// var slice = make([]int,4,5)
// // Size of slice can be increased by using Append()

// var newSlice = append(slice,40)

// //Map

// var idMap map[string]int

// func main() {
// 	// Variables
// 	welcone := "Welcome to go!" //short variable declearation

// 	println("Hello world")
// 	println(welcone)
// 	println(greet)
// 	println(pi)

// 	// If else
// 	if goLearner {
// 		println("Bravo")
// 	} else if !goLearner {
// 		println("please learn GO!")
// 	} else {
// 		println("Poor man!!")
// 	}

// 	// Switch

// 	switch {
// 	case goLearner:
// 		println("Brilliant")
// 	case !goLearner:
// 		println("Why not?")
// 	default:
// 		println("should learn Go")
// 	}

// 	switch os := runtime.GOOS; os {
// 	case "linus":
// 		println("Linux machine")

// 	default:
// 		println("maybe windows")
// 	}

// 	// arrays new style

// 	nums:= [...]int{2,3,5,6}
// 	fmt.Println(nums)

// 	// i - index --- v - curr value ---- nums arr
// 	for i, v := range nums {
// 		fmt.Printf("ind %d , val %d ",i,v) // print with assign
// 	}

// 	// slices
// s1 := vowels[1:3] // holds only starting and before ends index i.e --> i e
// s2 := vowels[2:5] // ==> e o u

// // also can create a slice b/c slices are underlying arr

// slice1 := []int {1,2,3,4}
// fmt.Printf("Slice %d",slice1)

// // methods of arr

// println(len(s1)) //2
// println(cap(s1)) // 4 because remaning elements are only 4.
// println(cap(s2)) //3 - because remaning elements are only 3.

// // Map
// idMap =make(map[string]int )

// // literal Map
// idMap := map[string]int {"joe":123}
// println(idMap)
// }

package main

type P struct {
x string
y int
}
func main() {
	// var n rune
	// var val int
	// var slice1 = make([]int, 0, 3)
	// val = cap(slice1)
	// for i := 0; i <= val; i++ {
	// 	if i >= len(slice1) {
	// 		val = val + 1
	// 	}
	// 	n = 0
	// 	fmt.Println("Enter a number to add: ")
	// 	fmt.Scan(&n)
	// 	if n == 0 {
	// 		fmt.Println("Exiting!")
	// 		break
	// 	}
	// 	slice1 = append(slice1, int(n))
	// }
	// sort.Ints(slice1)
	// fmt.Println(slice1)
	// x := [...]int{4,8,5}
	// y:= -1
	// 	for _, elt :=range x {
	// 		if elt > y {
	// 		y = elt
	// 		}
	// 	}
	// fmt.Print(y)
	// names := map[string]int {"ian":1,"harris":2}
	// for i,j:=range names {
	// if i == "harris" {
	// 	fmt.Print(i,j)
	// }
	// }
	

	
		
}

// func main() {
// 	slice := make([]int, 0, 3)

// 	for {
// 		fmt.Print("Enter an integer or 'X' to quit: ")
// 		var input string
// 		fmt.Scan(&input)

// 		if input == "X" || input == "x" {
// 			break
// 		}

// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter an integer or 'X' to quit.")
// 			continue
// 		}

// 		slice = append(slice, num)

// 		// Sort the slice
// 		sort.Ints(slice)

// 		fmt.Println("Sorted Slice:", slice)
// 	}
// }
