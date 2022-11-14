# SoT Editor (Speed of Thought Editor)

## Features


- [ ] Open smth under cursor (files, files on specific line and col pos, url)


_Current iteration:_
- https://viewsourcecode.org/snaptoken/kilo/02.enteringRawMode.html
- https://go.dev/tour/concurrency/1


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
    `nil` for all

Type conversion
:   `type(v)` - `float32(x)` converts **x** to float32 type
    > assignments between vars of different type requires an explicit conversion

Type assertion
:   ```go
    var x interface{} = 2
    isStr := x.(string) // -> panic
    // or
    val, isStr := x.(string) // -> nil false
    // but
    val, isInt := x.(int) // -> 2 true
    ```

    Flow control based on unknown type:
    ```go
    var x interface{} = 2
    swtich v := x.(type) {
    case int:
        // executes if x type is int
    case string:
        // executes if x type is string
    default:
        // default behaviour
    }
    ```
    > `x.(type)` can be used only in switch statement

### Type parameters

```go
                           x is the same type T like
                              any element of s
        type parameters         /
        ______|_______        _/_
func Max[T comparable](s []T, x T) int
            /          ¯¯|¯¯
           /             |
          /          slice of any type T
         /
   constraint for type T
```

Type parameters using to define generic type. Generic type could be useful for implementing generic data structures.

```go
type List[T any] struct {
    next *List[T]
    val T
}
```





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

Define a function (first class citizens)
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

Define closure
:   ```go
    func add(x int) func(int) int {
        return func(y int) int {
            return x + y
        }
    }
    ```

Define method-like function which is associated with specific type described as receier argument.
:   ```go
    type Square struct {
        x, y int
    }

    //  receiver argument
    //      |
    func (s Square) Area() int {
        return s.x * s.y
    }

    myvar := Square{2, 2}
    myvar.Area() // -> 4
    ```
    > it is possible declare a method with a receiver whose type is defined in the same package as the method.

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


### Pointer

Create
:   `var p *int` and then `p = &x`
    or
    `p := &x`

Read (i.e. dereferencing)
:   `fmt.Println(*p)`

Set value
:   `*p = 21` and now both `p` and `x` have value 21


### Array

Array's length can't be resized after defining.

Define
:   ```go
    var xs [6]int // var name [n]T
    // or with value
    xs := [6]int{2, 3, 5, 7, 11, 13}
    ```

### Slice

It's "dynamically-sized array" and pointer of an array's part

Define
:   ```go
    zs := []int{1,2,3,4}

    // or from array
    var zs []int = xs[1:4] // -> [3 5 7]
    // [x:y] - from `x` position (incl.) till `y` pos (not incl.)
    // [x:] - from `x` position (incl.) to the end
    // [:y] - from the beginning to `y` pos (incl.)
    // [:] - from the beginning to the end.
    ```

Length of a slice
:   ```go
    len(zs) // -> 3
    ```

    But slices also have capacity property. It's number of elements in pointed array (excluding number of elements before slice start position).
    ```go
    cap(zs) // -> 5
    ```

Extend
:   ```go
    zs = zs[:4] // -> [3 5 7 11]
    ```

Edit
:   ```go
    // append value
    append(zs, 22) // -> [3 5 7 11 22]
    ```


### Map

Maps keys to values

Define
:   ```go
    var kv = map[string]int {
        "age": 22,
        "maxAge": 67, // last comma is required!!!
    }
    ```

Put
:   ```go
    kv["age"] = 25
    ```

Get
:
    ```go
    a := kv["age"]
    ```

Delete
:   ```go
    delete(kv, "maxAge")
    ```

Check key existence
:   ```go
    a, has := kv["mykey"] // -> has is true if mykey is key of the kv
    ```



### Struct

It's collection of fields.

Create
:   ```go
    type Square struct {
        x, y int
    }

    z := Square{2, 2}
    ```

Read
:   ```go
    fmt.Println(z.x) -> 2
    fmt.Println(z) -> {2 2}
    ```

> `*` can be omitted in case of access to a struct by pointer


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
