# SoT Editor (Speed of Thought Editor)

## Features


- [ ] Open smth under cursor (files, files on specific line and col pos, url)


_Current iteration:_
- https://viewsourcecode.org/snaptoken/kilo/02.enteringRawMode.html
- https://go.dev/tour/moretypes/1


I learn Golang while creating console text-editor according to https://viewsourcecode.org/snaptoken/kilo

## Notes

Create a module
:   `go mod init <name>`

Run go app
:   `go run .`

Init app deps
:   `go tidy`

Redirect deps
:   `go mod edit -replace <module_name>=<location>`


### Types

Base types
:   bool
    string
    [u]int [u]int8/16/32/64 uintptr
    byte (alias for uint8)
    rune (alias for int32)
    float32/64
    complex64/128

Zero values
:   `0` for numeric types
    `false` for the boolean type
    `""` for strings

Type conversion
:   `type(v)` - `float32(x)` converts **x** to float32 type
    > assignments between vars of different type requires an explicit conversion



### Syntax

Declare a constant
:   `const X = 1`

Declare a variable
:   single: `var x int`
    multiple with the same type: `var x, y, z int`
    multiple: `var x int, y bool`
    with initialization: `var x, y int = 1, 2`

Declare and init a variable with implicit type (only inside a function)
:   `x := 1`

Define a function
:   ```
        name              result_type
         |                   |
    func Hello(name string) string { ... }
      |           |       \            |
     keyword      |        \         body - logic_here
              parameter    |
                     parameter_type
    ```
    > Functions started from capital letter are exported, other are private

Return multiple results from a function
:   ```
                          return types
                                |
    func Hello(name string) (string, int) {
      return name, 1
    }
    ```

Function named return
:   ```
    func Hello(name string) (x int) {
        x = 1
        return
    }          \
              return x as the function result
    ```

Loop
:   ```
         init statement (executed before the first iteration)
           |
           |      condition (evals before every iteration)
           |       /
    for i := 0; i < 10; i++ {
        ...               \
    }                      \
                         post statement (executed at the end of every iteration)
    ```
    > **init** and **post** are optional

    Short version
    ```
    for i < 10 {
        i += 1
    }
    ```

    Infinite loop
    ```
    for {
      ...
    }
    ```

Branching `if..else`
:   ```
      condition expression
         |
    if x < 0 {
        ...   ------ executed when condition expression is true
    } else { ... }
        |
     executed when condition expression is false
    ```

Branching `switch..case`
:   ```
            executed when val1 is
              matching with x
                  /
  comparable val /
           \    /
    switch x { /
      case val1:
          ...
      case valN:
          ...  \
       default: \
          ...    \
    }             \
            executed when val1 is not matched with x
              and valN is matching with x
    ```
    > A case may be any expression
    > `switch` without comparable value works like `switch true`

### Creating application

- Code executed an application must be placed in **main** package.
- In **main** package may be used `init` function to initialize app's settings


### What is module?

A module contains set of functions and types of specific domain.

```
        packageX --- code
        /
       /
Module ----- packageY --- code
      \
       \
      packageZ --- code
```

> module specifies own dependencies and even own go version.

> private functions in module are started with lower-case letter and upper-case for public using.
