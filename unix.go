package main
import"fmt"

func add(){

var num1,num2,sum int
fmt.Print("Enter the first number - ")
fmt.Scan(&num1)
fmt.Print("\nEnter the second number - ")
fmt.Scan(&num2)

sum=num1+num2
fmt.Println("\nsum of the two number is = ",sum)

}

func array(){
var sum int
arr:=[]int{1,2,3,4,5,6}
for i:=0;i<6;i++{
sum=sum+arr[i]
fmt.Print(arr[i])
}
fmt.Print("\n sum of the total array -",sum)

}
func main(){

//add()
array()
}

