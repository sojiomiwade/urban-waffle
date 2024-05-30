package main

import (
    "fmt"

    "oluwasoji.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    greetings.Helloworld()
    fmt.Println(message)
}
