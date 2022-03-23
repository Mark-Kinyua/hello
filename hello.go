package main

import (
    "fmt"
    "log"

    "github.com/Mark-Kinyua/greetings/src-d/enry/v2"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names
    names := []string{"Mark", "Ann", "Stella"}

    // Request a greeting messages.
    messages, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(messages)
}
