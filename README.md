# SoT Editor (Speed of Thought Editor)

## Features


- [ ] Open smth under cursor (files, files on specific line and col pos, url)


_Current iteration:_
- https://viewsourcecode.org/snaptoken/kilo/02.enteringRawMode.html
- https://go.dev/doc/tutorial/add-a-test


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


### Syntax

Declare a variable
:   `var x int`

Init a variable
:   ```
    var x int
    x = 1
    ```

Declare and init a variable
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
