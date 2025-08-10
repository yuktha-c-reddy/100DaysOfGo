// Day 1–2: Setup Go, variables, constants, fmt.

// Mini Project: Personal Finance Greeting CLI — takes name, monthly income, calculates saving goals.

package main
import "fmt"

func main() {
  fmt.Println("SIP Calculator")
  var name string
  var income float32
  var dur int
  var investment float32
  var minSavings float32
  fmt.Println("Enter your name and monthly income")
  fmt.Scanf("%s %f",&name , &income)
  if (name=="" || income<=0){
      fmt.Println("Please enter correct name and positive income")
  } else{
      minSavings=0.2*income
      var choice string
      fmt.Println("You have to save atleast 20% of your income = ",minSavings)
      fmt.Println("Do you want to change the savings amount? Y/N")
      
      fmt.Scanf("%s",&choice)
      if choice=="Y"{
        fmt.Println("Enter monthly investment and no.of months")
        fmt.Scanf("%d %f",&dur,&investment)
        if (dur<=0 || investment<0){
            fmt.Println("Please enter positive values")
        }else{
            total_invested:=(float32(dur)*investment)
            fmt.Printf("Total investment is %.2f \n" ,total_invested)
        }
      }else if choice=="N"{
         fmt.Println("Calculating savings per annum for ",minSavings)
         fmt.Println(minSavings*12.0)
      } else {
          fmt.Println("Please enter valid choice")
      }
      
    }
  
  }

