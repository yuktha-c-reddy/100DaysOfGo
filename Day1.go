// Day 1–2: Setup Go, variables, constants, fmt.

// Mini Project: Personal Finance Greeting CLI — takes name, monthly income, prints a welcome message with saving goals.

package main
import "fmt" //formatted i/p & o/p

func main(){
    fmt.Println("Personal Finance Greeting CLI")
    fmt.Println("Enter name and monthly income")
    var income float32
    var name string
    fmt.Scanf("%s %f",&name,&income)
    fmt.Printf("Dear %s , Kindly save at least %.2f rupees per month\n",name, 0.2*income)   
}


