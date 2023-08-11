package main

import (
	"fmt"
)

// concept of Expression switch
// statement
func switch1() {
	switch day := 7; day {
	case 1:
		fmt.Println("Monday")
		break
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Choice !")
	}
}

// concept of Expression switch
// statement
func switch2() {
	var value int = 2
	switch {
	case value == 1:
		fmt.Println("Hello India")
	case value == 2:
		fmt.Println("Hello Chandu")
	case value == 3:
		fmt.Println("Hello World")
	default:
		fmt.Println("Invalid Choice !")

	}
}

// concept of Expression switch
// statement
func switch3() {
	var value string = "one"
	switch value {
	case "one":
		fmt.Println("Chandu kumar")
	case "two":
		fmt.Println("Go language")
	case "three":
		fmt.Println("Github")
	default:
		fmt.Println("Invalid Choice")

	}
}

// concept of Type switch
// statement
func switch4() {
	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value of boolean type")
	case int:
		fmt.Println("value of Integer type")
	case float64:
		fmt.Println("value of float64 type")
	case string:
		fmt.Println("value of string type")
	default:
		fmt.Printf("value is of Default %T", q)
	}
}

// concept of Expression switch
// statement
func switch5() {
	var value string = "one"
	// switch statement without default statement multiple value in case statements
	switch value {
	case "one":
		fmt.Println("C#")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}
}
func main() {

	switch5()
	//switch4()
	//switch3()
	//switch2()
	//switch1()
}
