
# Composite Data Types in Go

This directory provides illustrative examples of Go's composite data types, including arrays, slices, maps, structs, and functions. Each example demonstrates the declaration, initialization, and manipulation of these types, offering insights into their usage and behavior.

---

## 1. Arrays

Arrays in Go are fixed-size sequences of elements with the same type. They are value types, meaning assigning one array to another copies all elements.

**Key Concepts:**

- **Declaration and Initialization:**
  ```go
  var numbers [5]int
  primes := [3]int{2, 3, 5}
  ```

- **Accessing Elements:**
  ```go
  first := numbers[0]
  ```

- **Iterating Over Arrays:**
  ```go
  for i, v := range primes {
      fmt.Println(i, v)
  }
  ```

- **Copying Arrays:**
  ```go
  var copy [3]int
  copy = primes
  ```

---

## 2. Slices

Slices are dynamically-sized, flexible views into arrays. They are more commonly used than arrays due to their versatility.

**Key Concepts:**

- **Declaration and Initialization:**
  ```go
  var s []int
  s = []int{1, 2, 3}
  ```

- **Creating Slices with `make`:**
  ```go
  s := make([]int, 5, 10)
  ```

- **Appending Elements:**
  ```go
  s = append(s, 4, 5)
  ```

- **Slicing Arrays:**
  ```go
  arr := [5]int{1, 2, 3, 4, 5}
  s := arr[1:4] // [2 3 4]
  ```

- **Copying Slices:**
  ```go
  t := make([]int, len(s))
  copy(t, s)
  ```

- **Length and Capacity:**
  ```go
  len(s)
  cap(s)
  ```

---

## 3. Maps

Maps are unordered collections of key-value pairs. Keys must be of a type that is comparable (e.g., integers, strings).

**Key Concepts:**

- **Declaration and Initialization:**
  ```go
  m := make(map[string]int)
  m["one"] = 1
  ```

- **Accessing Values:**
  ```go
  value := m["one"]
  ```

- **Checking for Existence:**
  ```go
  val, ok := m["two"]
  ```

- **Deleting Entries:**
  ```go
  delete(m, "one")
  ```

- **Iterating Over Maps:**
  ```go
  for k, v := range m {
      fmt.Println(k, v)
  }
  ```

---

## 4. Structs

Structs are composite data types that group together variables under a single type. They are useful for representing complex data structures.

**Key Concepts:**

- **Defining Structs:**
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

- **Initializing Structs:**
  ```go
  p := Person{Name: "Alice", Age: 30}
  ```

- **Accessing Fields:**
  ```go
  fmt.Println(p.Name)
  ```

- **Pointers to Structs:**
  ```go
  pp := &p
  pp.Age = 31
  ```

- **Embedding Structs:**
  ```go
  type Employee struct {
      Person
      EmployeeID int
  }
  ```

---

## 5. Functions

Functions are first-class citizens in Go, meaning they can be assigned to variables, passed as arguments, and returned from other functions.

**Key Concepts:**

- **Defining Functions:**
  ```go
  func add(a int, b int) int {
      return a + b
  }
  ```

- **Multiple Return Values:**
  ```go
  func divide(a, b int) (int, error) {
      if b == 0 {
          return 0, errors.New("division by zero")
      }
      return a / b, nil
  }
  ```

- **Anonymous Functions:**
  ```go
  func() {
      fmt.Println("Hello, World!")
  }()
  ```

- **Function Variables:**
  ```go
  var f func(int, int) int
  f = add
  ```

- **Passing Functions as Arguments:**
  ```go
  func operate(a int, b int, op func(int, int) int) int {
      return op(a, b)
  }
  ```

---

## Conclusion

This collection of examples serves as a practical guide to understanding and utilizing Go's composite data types. By studying these examples, you can gain a deeper insight into how arrays, slices, maps, structs, and functions operate in Go, enabling you to write more efficient and idiomatic Go code.
