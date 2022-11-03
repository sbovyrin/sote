package main

import (
    "fmt"
    "sb/greetings"
)

func main() {
    message := greetings.Hello("Sergey")
    fmt.Println(message)
}
