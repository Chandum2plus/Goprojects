package main

import (
	"fmt"
	"os"
)

// Example of the command line argument

// Commands -
//     ~/Desktop/Goproj/src/Goprojects/codde    Chandu !1 ?9 ─ go build CommmandLineArgs.go
//     ~/Desktop/Goproj/src/Goprojects/codde    Chandu !1 ?9 ─ go run CommandLineArgs.go Chandu kumar    "then Press enter"

// step and Commands for run the code
/*
❯ go build CommandLineArgs.go
❯ go run CommandLineArgs.go Chandu kumar Paras kumar sonal kumar Saurabh kumar
[/var/folders/51/6nj0zf6j6_3759yqlg2cl2m80000gn/T/go-build2022142989/b001/exe/CommandLineArgs Chandu kumar Paras kumar sonal kumar Saurabh kumar]
[Chandu kumar Paras kumar sonal kumar Saurabh kumar]
[/var/folders/51/6nj0zf6j6_3759yqlg2cl2m80000gn/T/go-build2022142989/b001/exe/CommandLineArgs Chandu kumar Paras kumar sonal kumar Saurabh kumar]

*/
func main() {
	argWithcommand := os.Args
	argWithoutCommand := os.Args[1:]
	arg := os.Args
	fmt.Println(argWithcommand)
	fmt.Println(argWithoutCommand)
	fmt.Println(arg)
}
