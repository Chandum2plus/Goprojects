package main
import "fmt"

// Array of rune
func Rune(){
var name=[6]rune{'c','h','a','n','d','u'}
for i:=0;i<len(name);i++{
fmt.Printf("%c\n",name[i])

}

//  Example of Rune Array using user input
func UserRune(name string){
//name:=string
arr:=[100]rune{}
for i,v:=range name{
arr[i]=v
fmt.Println("rune =",string(arr[i]))
}

}
func main(){
//arr:="Chandu"
var name= string
fmt.Print("Enter your name -")
fmt.Scan(&name)
UserRune(name)

}

