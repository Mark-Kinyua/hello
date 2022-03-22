package main

import (
    "fmt"

    "github.com/Mark-Kinyua/greetings/src-d/enry/v2"
)

func main() {
    // Get a greeting message and print it.
    message1 := greetings.Hello("Welcome to the program")
    fmt.Println(message1)
}
