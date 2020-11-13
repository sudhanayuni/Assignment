package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp5 := Employee{
        firstName: "John",
        lastName:  "Paul",
    }
    fmt.Println("First Name:", emp5.firstName)
    fmt.Println("Last Name:", emp5.lastName)
    fmt.Println("Age:", emp5.age)
    fmt.Println("Salary:", emp5.salary)
}