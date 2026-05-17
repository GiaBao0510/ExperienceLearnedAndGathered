# Functions trong Go - Hướng dẫn Đầy đủ

## 📋 Mục lục

1. [Function cơ bản](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#1-function-c%C6%A1-b%E1%BA%A3n)
2. [Parameters và Return Values](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#2-parameters-v%C3%A0-return-values)
3. [Named Return Values](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#3-named-return-values)
4. [Variadic Functions](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#4-variadic-functions)
5. [Anonymous Functions](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#5-anonymous-functions)
6. [Defer Statement](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#6-defer-statement)
7. [Functions với Slice](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#7-functions-v%E1%BB%9Bi-slice)
8. [Higher-order Functions](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#8-higher-order-functions)
9. [Best Practices](https://claude.ai/chat/15d8c184-b0e9-425c-a108-2671bce677f7#9-best-practices)

---

## 1. Function cơ bản

### 1.1. Function là gì?

**Function (hàm)** là một khối code độc lập thực hiện một nhiệm vụ cụ thể.

**Lợi ích:**

- ✅ Tái sử dụng code
- ✅ Code dễ đọc, dễ maintain
- ✅ Tách biệt logic
- ✅ Dễ test

### 1.2. Cú pháp

**Function không return:**

```go
func functionName(param1 type1, param2 type2) {
    // Code
}
```

**Function có return:**

```go
func functionName(param1 type1, param2 type2) returnType {
    // Code
    return value
}
```

### 1.3. Ví dụ cơ bản

```go
package main

import "fmt"

// Function không return
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Function có return
func add(a int, b int) int {
    return a + b
}

// Shorthand - cùng type
func multiply(a, b int) int {  // a int, b int → a, b int
    return a * b
}

func main() {
    greet("John")            // Hello, John
    
    sum := add(5, 3)
    fmt.Println("Sum:", sum) // Sum: 8
    
    product := multiply(4, 5)
    fmt.Println("Product:", product)  // Product: 20
}
```

---

## 2. Parameters và Return Values

### 2.1. Multiple Parameters

```go
package main

import "fmt"

// Nhiều parameters cùng type
func fullName(firstName, lastName string) string {
    return firstName + " " + lastName
}

// Nhiều parameters khác type
func printInfo(name string, age int, salary float64) {
    fmt.Printf("Name: %s, Age: %d, Salary: %.2f\n", name, age, salary)
}

func main() {
    name := fullName("John", "Doe")
    fmt.Println(name)  // John Doe
    
    printInfo("Alice", 25, 50000.50)
    // Name: Alice, Age: 25, Salary: 50000.50
}
```

### 2.2. Multiple Return Values

**Go hỗ trợ return nhiều giá trị - feature rất hữu ích!**

```go
package main

import "fmt"

// Return 2 giá trị
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Return value và error (pattern phổ biến)
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    // Nhận cả 2 giá trị
    q, r := divide(17, 5)
    fmt.Printf("17 / 5 = %d remainder %d\n", q, r)
    // 17 / 5 = 3 remainder 2
    
    // Bỏ qua giá trị không cần với _
    q2, _ := divide(20, 3)
    fmt.Println("Quotient:", q2)  // Quotient: 6
    
    // Error handling
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)  // Error: cannot divide by zero
    } else {
        fmt.Println("Result:", result)
    }
}
```

### 2.3. Không có Default Parameters

**⚠️ Go KHÔNG hỗ trợ default parameters!**

```go
// ❌ COMPILE ERROR
func greet(name string = "Guest") {
    fmt.Println("Hello,", name)
}
```

**✅ Workaround - Check inside function:**

```go
package main

import "fmt"

func greet(name string) {
    if name == "" {
        name = "Guest"  // Default value
    }
    fmt.Println("Hello,", name)
}

func main() {
    greet("John")   // Hello, John
    greet("")       // Hello, Guest
}
```

**✅ Better - Variadic function hoặc Options pattern:**

```go
// Option 1: Variadic
func greet(names ...string) {
    if len(names) == 0 {
        fmt.Println("Hello, Guest")
        return
    }
    for _, name := range names {
        fmt.Println("Hello,", name)
    }
}

// Option 2: Options struct
type GreetOptions struct {
    Name     string
    Greeting string
}

func greetWithOptions(opts GreetOptions) {
    greeting := opts.Greeting
    if greeting == "" {
        greeting = "Hello"
    }
    
    name := opts.Name
    if name == "" {
        name = "Guest"
    }
    
    fmt.Printf("%s, %s\n", greeting, name)
}

func main() {
    greet()              // Hello, Guest
    greet("John", "Jane") // Hello, John \n Hello, Jane
    
    greetWithOptions(GreetOptions{})  // Hello, Guest
    greetWithOptions(GreetOptions{Name: "John"})  // Hello, John
    greetWithOptions(GreetOptions{Name: "Jane", Greeting: "Hi"})  // Hi, Jane
}
```

---

## 3. Named Return Values

### 3.1. Cú pháp

```go
func functionName(params) (returnName1 type1, returnName2 type2) {
    // returnName1 và returnName2 đã được khai báo
    returnName1 = value1
    returnName2 = value2
    return  // Naked return
}
```

### 3.2. Ví dụ

```go
package main

import "fmt"

// Named return values
func divide(a, b int) (quotient int, remainder int) {
    quotient = a / b
    remainder = a % b
    return  // Naked return - tự return quotient, remainder
}

// So sánh với unnamed
func divideUnnamed(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder  // Phải chỉ rõ
}

// Named return với early return
func safeDivide(a, b int) (result int, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return  // result = 0 (zero value), err = error
    }
    result = a / b
    return  // result = a/b, err = nil (zero value)
}

func main() {
    q, r := divide(17, 5)
    fmt.Printf("%d, %d\n", q, r)  // 3, 2
    
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println(err)  // division by zero
    } else {
        fmt.Println(result)
    }
}
```

### 3.3. Lưu ý với Named Returns

**⚠️ Named return có thể bị "shadow":**

```go
func tricky() (result int) {
    result = 10
    
    // ❌ Tạo biến mới result, shadow outer result
    result := 20
    fmt.Println(result)  // 20
    
    return  // Return outer result = 10, không phải 20!
}

func main() {
    fmt.Println(tricky())  // 10, không phải 20!
}
```

**✅ Fix:**

```go
func fixed() (result int) {
    result = 10
    
    // ✅ Dùng = thay vì :=
    result = 20
    fmt.Println(result)  // 20
    
    return  // Return 20
}
```

---

## 4. Variadic Functions

### 4.1. Variadic là gì?

**Variadic function** nhận **số lượng tham số thay đổi** cùng kiểu.

**Cú pháp:**

```go
func functionName(params ...type) {
    // params là slice
}
```

### 4.2. Ví dụ cơ bản

```go
package main

import "fmt"

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum())           // 0
    fmt.Println(sum(1))          // 1
    fmt.Println(sum(1, 2))       // 3
    fmt.Println(sum(1, 2, 3, 4)) // 10
    
    // Spread slice
    nums := []int{5, 6, 7}
    fmt.Println(sum(nums...))    // 18
}
```

### 4.3. Variadic phải ở cuối

```go
// ✅ OK
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s\n", greeting, name)
    }
}

// ❌ COMPILE ERROR - variadic phải là param cuối
// func wrong(names ...string, greeting string) { }

func main() {
    greet("Hello", "John", "Jane", "Bob")
    // Hello, John
    // Hello, Jane
    // Hello, Bob
}
```

### 4.4. Variadic với interface{}

```go
package main

import "fmt"

// Printf-style function
func printf(format string, args ...interface{}) {
    fmt.Printf(format, args...)
}

func main() {
    printf("Name: %s, Age: %d, Salary: %.2f\n", "John", 25, 50000.50)
    // Name: John, Age: 25, Salary: 50000.50
}
```

### 4.5. Unpack slice

```go
package main

import "fmt"

func print(args ...interface{}) {
    fmt.Println(args...)
}

func main() {
    var a = []interface{}{123, "abc"}
    
    // Unpack - truyền từng element
    print(a...)  // 123 abc
    
    // Không unpack - truyền slice
    print(a)     // [123 abc]
}
```

---

## 5. Anonymous Functions

### 5.1. Anonymous Function là gì?

**Anonymous function** (hàm ẩn danh) là hàm **không có tên**.

### 5.2. Cú pháp

```go
// Named function
func add(a, b int) int {
    return a + b
}

// Anonymous function
var add = func(a, b int) int {
    return a + b
}
```

### 5.3. Ví dụ

```go
package main

import "fmt"

func main() {
    // Define và gọi ngay
    func() {
        fmt.Println("Hello from anonymous function")
    }()
    
    // Gán vào biến
    add := func(a, b int) int {
        return a + b
    }
    
    result := add(5, 3)
    fmt.Println("Result:", result)  // Result: 8
    
    // Closure - access outer variables
    x := 10
    increment := func() {
        x++
        fmt.Println("x:", x)
    }
    
    increment()  // x: 11
    increment()  // x: 12
    fmt.Println("Final x:", x)  // Final x: 12
}
```

### 5.4. Closures

```go
package main

import "fmt"

// Function trả về function
func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    counter1 := makeCounter()
    counter2 := makeCounter()
    
    fmt.Println(counter1())  // 1
    fmt.Println(counter1())  // 2
    fmt.Println(counter1())  // 3
    
    fmt.Println(counter2())  // 1 - independent
    fmt.Println(counter2())  // 2
}
```

---

## 6. Defer Statement

### 6.1. Defer là gì?

**`defer`** trì hoãn việc thực thi function cho đến khi **function bao ngoài return**.

### 6.2. Cú pháp

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}

// Output:
// hello
// world
```

### 6.3. Defer Stack (LIFO)

**Defer được push vào stack → thực thi theo thứ tự ngược lại (LIFO).**

```go
package main

import "fmt"

func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    
    fmt.Println("Start")
}

// Output:
// Start
// 3
// 2
// 1
```

### 6.4. Use Cases

#### **Use Case 1: Close Resources**

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Create("test.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()  // Đảm bảo close file
    
    // Write to file
    fmt.Fprintln(f, "Hello, World!")
    
    // File sẽ tự close khi main() return
}
```

#### **Use Case 2: Unlock Mutex**

```go
import "sync"

var mu sync.Mutex

func criticalSection() {
    mu.Lock()
    defer mu.Unlock()  // Đảm bảo unlock
    
    // Critical code
    // Nếu panic, vẫn unlock
}
```

#### **Use Case 3: Recover from Panic**

```go
package main

import "fmt"

func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    
    panic("Something went wrong!")
}

func main() {
    safeFunction()
    fmt.Println("Program continues")
}

// Output:
// Recovered from: Something went wrong!
// Program continues
```

### 6.5. Defer và Return Values

**Defer có thể modify named return values:**

```go
package main

import "fmt"

func yes() (result string) {
    defer func() {
        result = "no"  // Modify return value
    }()
    return "yes"
}

func main() {
    fmt.Println(yes())  // no, không phải yes!
}
```

**Giải thích:**

```
1. return "yes" → result = "yes"
2. Defer execute → result = "no"
3. Function return → "no"
```

### 6.6. Defer Arguments Evaluated Immediately

```go
package main

import "fmt"

func main() {
    x := 10
    
    defer fmt.Println("Deferred:", x)  // x = 10 được evaluate ngay
    
    x = 20
    
    fmt.Println("Current:", x)  // 20
}

// Output:
// Current: 20
// Deferred: 10  ← 10, không phải 20!
```

**Workaround - Dùng closure:**

```go
func main() {
    x := 10
    
    defer func() {
        fmt.Println("Deferred:", x)  // Closure - refer to x
    }()
    
    x = 20
    
    fmt.Println("Current:", x)
}

// Output:
// Current: 20
// Deferred: 20  ← 20!
```

---

## 7. Functions với Slice

### 7.1. Pass by Value vs Reference

![Slice Structure](https://zalopay-oss.github.io/go-advanced/images/slice_1.png)

**Go là pass by value, nhưng slice header chứa pointer → có thể modify underlying array.**

### 7.2. Array vs Slice

```go
package main

import "fmt"

// Array - pass by value (copy toàn bộ)
func modifyArray(arr [3]int) {
    arr[0] = 100
    fmt.Println("Inside modifyArray:", arr)
}

// Slice - pass slice header (có pointer)
func modifySlice(s []int) {
    s[0] = 100
    fmt.Println("Inside modifySlice:", s)
}

func main() {
    // Array
    arr := [3]int{1, 2, 3}
    modifyArray(arr)
    fmt.Println("After modifyArray:", arr)  // [1 2 3] - không đổi!
    
    fmt.Println()
    
    // Slice
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println("After modifySlice:", slice)  // [100 2 3] - đã đổi!
}

// Output:
// Inside modifyArray: [100 2 3]
// After modifyArray: [1 2 3]
//
// Inside modifySlice: [100 2 3]
// After modifySlice: [100 2 3]
```

### 7.3. Append trong Function

**⚠️ Lưu ý: `append` có thể tạo slice mới!**

```go
package main

import "fmt"

func appendValue(s []int, val int) {
    s = append(s, val)
    fmt.Println("Inside:", s)
}

func main() {
    slice := []int{1, 2, 3}
    appendValue(slice, 4)
    fmt.Println("Outside:", slice)  // [1 2 3] - không đổi!
}

// Output:
// Inside: [1 2 3 4]
// Outside: [1 2 3]
```

**✅ Fix - Return slice:**

```go
func appendValue(s []int, val int) []int {
    return append(s, val)
}

func main() {
    slice := []int{1, 2, 3}
    slice = appendValue(slice, 4)
    fmt.Println("Outside:", slice)  // [1 2 3 4]
}
```

**✅ Fix - Pointer to slice:**

```go
func appendValue(s *[]int, val int) {
    *s = append(*s, val)
}

func main() {
    slice := []int{1, 2, 3}
    appendValue(&slice, 4)
    fmt.Println("Outside:", slice)  // [1 2 3 4]
}
```

---

## 8. Higher-order Functions

### 8.1. Function as Parameter

```go
package main

import "fmt"

// Function nhận function làm parameter
func apply(nums []int, fn func(int) int) []int {
    result := make([]int, len(nums))
    for i, num := range nums {
        result[i] = fn(num)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    
    // Double
    doubled := apply(nums, func(x int) int {
        return x * 2
    })
    fmt.Println("Doubled:", doubled)  // [2 4 6 8 10]
    
    // Square
    squared := apply(nums, func(x int) int {
        return x * x
    })
    fmt.Println("Squared:", squared)  // [1 4 9 16 25]
}
```

### 8.2. Function as Return Value

```go
package main

import "fmt"

// Function trả về function
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

---

## 9. Best Practices

### 9.1. Function Naming

```go
// ✅ Good - Verb + Noun
func calculateTotal() {}
func getUserById() {}
func validateEmail() {}

// ❌ Bad - Vague
func process() {}
func handle() {}
func do() {}
```

### 9.2. Function Size

```go
// ✅ Good - Nhỏ, single responsibility
func validateEmail(email string) bool {
    // Validation logic
}

func saveUser(user User) error {
    // Save logic
}

// ❌ Bad - Quá dài, nhiều responsibility
func processUser(email string, name string) error {
    // Validate email
    // Validate name
    // Create user
    // Save to DB
    // Send email
    // Log activity
    // Update cache
    // ... 100 lines
}
```

### 9.3. Error Handling

```go
// ✅ Good - Return error
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// ❌ Bad - Panic
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")  // Avoid panic
    }
    return a / b
}
```

### 9.4. Named Returns

```go
// ✅ Good - Short, clear
func divide(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return
}

// ⚠️ Caution - Long function với named return dễ confuse
func longFunction() (result int, err error) {
    // ... 50 lines
    result = calculate()
    // ... 50 lines
    return  // Khó theo dõi giá trị return
}
```

### 9.5. Defer Usage

```go
// ✅ Good
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()  // Ngay sau khi open
    
    // Read file
    return nil
}

// ❌ Bad
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    
    // ... nhiều code
    
    defer f.Close()  // Quá xa, dễ quên
    return nil
}
```

---

## 📚 Tổng kết

### Function Basics

```go
// Simple
func add(a, b int) int {
    return a + b
}

// Multiple returns
func divide(a, b int) (int, int) {
    return a / b, a % b
}

// Named returns
func divide(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return
}

// Variadic
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### Defer

```go
// LIFO order
defer close()
defer unlock()
defer cleanup()

// With resources
f, _ := os.Open("file")
defer f.Close()

// With panic recovery
defer func() {
    if r := recover(); r != nil {
        log.Println("Recovered:", r)
    }
}()
```

### Best Practices

1. ✅ Small functions (single responsibility)
2. ✅ Clear naming (verb + noun)
3. ✅ Return errors, don't panic
4. ✅ Defer for cleanup
5. ✅ Named returns cho short functions
6. ❌ Không dùng named returns cho long functions
7. ❌ Không panic trong library code

---

_Functions là building blocks của mọi chương trình!_