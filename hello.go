package main

import (
    "fmt"
    // "log"
    // "sb/greetings"
)

func main() {
    var x interface{} = 2
    isStr := x.(string) // -> panic
    val, isStr := x.(string) // -> nil false
    val, isInt := x.(int) // -> 2 true

    // log.SetPrefix("greetings: ")
    // log.SetFlags(0)

    // names := []string{"Sergey", "John", "Jane"}
    // messages, err := greetings.Hellos(names)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(messages)
}
