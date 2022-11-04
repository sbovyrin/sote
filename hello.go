package main

import (
    "fmt"
    "log"
    "sb/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Sergey", "John", "Jane"}
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}
