package main

import "fmt"

const const_str = "const test" //constant variable

func plus(a int, b int) int { // function
	return a+b
}

func vals() (int, int) { //function return mutiple value
    return 3, 7
}

func main(){
	fmt.Println("hello world") // print to console
	
	fmt.Println("7.0/3.0 =", 7.0/3.0, "value") // value in go lang

	var str string = "test variable" // declaring and initializing a variable

	str1 := "test variable 1" // shorthand for declaring and initializing a variable variable is using ":=" operator.

	fmt.Println(str, str1)
	str  = "reassign variable"
	str1 = "reassign variable 1"
	fmt.Println(str, str1)

	for j := 1; j <= 9; j++ { //for and if/else syntax
        if j%2 == 0 {
        	fmt.Println(j, "la so chan")
        } else {
        	fmt.Println(j, "la so le")
        }
    }

    i := 2
    fmt.Print("write ", i, " as ") // switch/case syntax
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    var a [5]int // array
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("len:", len(a))

    b := [2][3]int{{1,2,3}, {4,5,6}}
    fmt.Println("emp:", b)

    res := plus(1, 2)
    fmt.Println("1 + 2 = ", res)

    x, z := vals()
    fmt.Println(x,z)
}
