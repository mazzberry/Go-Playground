
# 🧮 Basic Operations in Go

This folder contains a collection of simple and educational examples to help you understand the basic concepts of the Go programming language. Each file demonstrates a fundamental topic, making it ideal for beginners.

## 📁 Files with Examples

### 📌 variables.go

Introduction to declaring and using variables in Go:

```go
package main

import "fmt"

func main() {
    var name string = "Alice"
    age := 25
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

---

### 📌 constants.go

Defining and using constants:

```go
package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "Earth"
    fmt.Println("Pi:", Pi)
    fmt.Println("Planet:", World)
}
```

---

### 📌 operators.go

Demonstration of different types of operators:

```go
package main

import "fmt"

func main() {
    a, b := 10, 3
    fmt.Println("Addition:", a+b)
    fmt.Println("Subtraction:", a-b)
    fmt.Println("Multiplication:", a*b)
    fmt.Println("Division:", a/b)
    fmt.Println("Modulus:", a%b)

    fmt.Println("Equal:", a == b)
    fmt.Println("Not Equal:", a != b)
    fmt.Println("Greater Than:", a > b)
    fmt.Println("Logical AND:", a > 1 && b < 5)
}
```

---

### 📌 input_output.go

Taking user input and printing output:

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}
```

---

### 📌 data_types.go

Overview of basic data types:

```go
package main

import "fmt"

func main() {
    var str string = "Go"
    var num int = 42
    var decimal float64 = 3.14
    var truth bool = true

    fmt.Println("String:", str)
    fmt.Println("Integer:", num)
    fmt.Println("Float:", decimal)
    fmt.Println("Boolean:", truth)
}
```

---

### 📌 type_conversion.go

Type casting and conversion:

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)

    fmt.Println("Integer:", i)
    fmt.Println("Float:", f)
    fmt.Println("Unsigned:", u)
}
```

---

### 📌 control_structures.go

Control flow structures like if, switch, and loops:

```go
package main

import "fmt"

func main() {
    x := 5

    if x > 0 {
        fmt.Println("Positive")
    } else {
        fmt.Println("Non-positive")
    }

    switch x {
    case 1:
        fmt.Println("One")
    case 5:
        fmt.Println("Five")
    default:
        fmt.Println("Other number")
    }

    for i := 0; i < 3; i++ {
        fmt.Println("Iteration:", i)
    }
}
```

---

## 🚀 How to Run

To run any of the files:

1. Make sure Go is installed on your system:

```bash
go version
```

2. Run a file using the following command:

```bash
go run <filename>.go
```

Example:

```bash
go run variables.go
```

## 🎯 Purpose

The goal of this collection is to provide simple examples for a better understanding of basic programming concepts in Go. These examples can serve as a starting point for deeper learning and exploration of the language.

---
