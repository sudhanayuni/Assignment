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

    //creating struct specifying field names
    emp1 := Employee{
        firstName: "Sudha",
        age:       22,
        salary:    500000,
        lastName:  "nayuni",
    }

    //creating struct without specifying field names
    emp2 := Employee{"Thomas", "Paul", 29, 800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}