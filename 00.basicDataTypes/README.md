
# Basic Data Types in Go

This directory provides illustrative examples of Go's basic data types, including booleans, numeric types, strings, and characters. Each example demonstrates the declaration, initialization, and usage of these types, offering insights into their behavior and applications.

---

## 1. Boolean Type

The `bool` type represents boolean values and can hold one of two constants: `true` or `false`. It is commonly used in conditional statements and logical operations.

**Example:**

```go
package main

import "fmt"

func main() {
    var isActive bool = true
    fmt.Println("Is Active:", isActive)
}
```

---

## 2. Numeric Types

Go provides various numeric types to represent integer and floating-point numbers. These include:

- **Integers:**
  - Signed: `int`, `int8`, `int16`, `int32`, `int64`
  - Unsigned: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Floating-Point:**
  - `float32`, `float64`
- **Complex Numbers:**
  - `complex64`, `complex128`

**Example:**

```go
package main

import "fmt"

func main() {
    var age int = 30
    var salary float64 = 50000.50
    var complexNum complex128 = complex(2, 3)

    fmt.Println("Age:", age)
    fmt.Println("Salary:", salary)
    fmt.Println("Complex Number:", complexNum)
}
```

---

## 3. String Type

The `string` type represents a sequence of characters. Strings in Go are immutable, meaning once created, their contents cannot be changed.

**Example:**

```go
package main

import "fmt"

func main() {
    var greeting string = "Hello, World!"
    fmt.Println(greeting)
}
```

---

## 4. Character Representation

Go does not have a dedicated `char` type. Instead, characters are represented using the `rune` type, which is an alias for `int32`. This allows for the representation of Unicode code points.

**Example:**

```go
package main

import "fmt"

func main() {
    var letter rune = 'A'
    fmt.Printf("Character: %c, Unicode: %U\n", letter, letter)
}
```

---

## 5. Type Conversion

Go requires explicit type conversion when assigning values between different types. This ensures type safety and prevents unintended data loss.

**Example:**

```go
package main

import "fmt"

func main() {
    var x int = 10
    var y float64 = float64(x)
    fmt.Println("Integer:", x)
    fmt.Println("Converted to Float:", y)
}
```

---

## Conclusion

Understanding Go's basic data types is fundamental to writing effective Go programs. This directory's examples provide a practical overview of how to declare, initialize, and use these types in various scenarios.

For more detailed information, refer to the official Go documentation: [https://go.dev/ref/spec#Types](https://go.dev/ref/spec#Types)
