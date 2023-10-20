package main

import "fmt"

// find the prime number

type Prime interface {
	PrimeNumber(num int)
}

// this  is the variable created by the type keyword
type primenum int

func (p primenum) PrimeNumber(num int) {
	i := 0
	Primecount := 0
	for i = 2; i < num/2; i++ {
		Primecount++
		break
	}
	if Primecount == 0 && Primecount != 1 {
		fmt.Println(num, "-Prime Number")
	} else {
		fmt.Println(num, " - Not prime Number ")
	}

}

func main() {
	var num Prime
	num = primenum(0)
	//fmt.Println("Enter the Number - ")
	//fmt.Scan(&num)
	//num.PrimeNumber(3)
	num.PrimeNumber(23)

}
