# GO Laboratory

## Foundations of Golang

### Identifiers, Filenames, and Keywords

Identifiers in Go may be names of functions, constants, variables, statement 
labels, types, or package names.

Go identifiers must start with a letter or underscore, not a number.  The
identifier itself can contain letters, numbers, and underscores.

- Identifiers are case sensitive.
- It's best to keep identifiers between 4 and 15 characters in length.

The underscore is a **black identifier** and is used as an anonymous 
placeholder with a special meaning in certain parts of the program.

Go also has some identifiers that have been prepared and can be used 
for constants, functions, and types.

- Constants: true, false, itoa, nil
- Functions: make, copy, append, close, delete, plus others
- Types: int, uint, string, bool, error, plus others

*Filenames*, which are used to identify each file uniquely, are typically 
written in lowercase and multiple words are separated with an underscore.

*Keywords* already have a meaning and therefore can't be used for 
identifiers.  There are 25 keywrods in Go, including map, struct, goto, 
switch, import, and return.

### Functions

*Functions* are blocks of code that perform a specific task.

In Go, every program has at least one function: main().  
This is a function declaration, which tells the program the name of the 
function, the return type, parameters, and the *definition*.  The definition
is the set of instructions or the function body.

There are other standard functions in Go such as len(), which will give you
the length of the argument, string, or array, depending on what was called.

This is the syntax of a function in Go:

```Golang
package example 

func function_name([parameter list]) [return_types] {
	// body of the function
}
```

- func starts the declaration, telling the program that a function is coming up
- function_name names the function being used
- parameter is similar to a placeholder, and gives the type, order, and number of parameters; there may be no parameters for a function 
- return_types is the list of data types for the values the function returns; this isn't always required
- body defines what the function does

Examples

```Golang
package example

func max(num1, num2 int) int {
	var result int // this is a local variable declaration.
	
	if (num1 > num2) {
		result = num1
    } else {
		result = num2
    }
	
	return result
}
```

### Data Types

The type of a variable specifies how much space it occupies and how the bit 
pattern should be interpreted.  In go there are four categories of data types:

1. Boolean
2. Numeric
3. Strings
4. Derived Types

*Bool* is one of the four elementary types.

Numeric types can be either integers, or floating-point values.  Exampels of 
numeric types include:

- uint8 (unsigned 8-bit integers from 0 to 255) 
- int16 (signed 16-bit integers)
- and the like for integers
- float32 (IEEE 32-bit floating-point)
- float64 (IEEE 64-bit floating-point)
- complex62 (float32 with real and imaginary parts)
- complex128 (float64 with real and imaginary parts)

Both *int* and *float* are also elementary types

Strings represent the value of the string, which is a sequence of bytes.  Once
created, the string cannot be changed.  Strings can be alphanumeric.  A string
is also an elementary type.

Derived types include pointer, array , structure , map types, etc..

### Constants

Note: Go doesn't allow you to mix numeric types like C and C++.  There is no 
automated type conversion.

An untyped constant that's an integer can only be used where integers are 
allowed.  You can't assign an integer constant to a string or boolean variable.
An untyped floating-point constant can be used wherever a floating point is
permitted.

All untyped constants in Go have default types.  constants that are integers 
default to "int", floats default to "float64", characters to "rune", etc.  
When you declare a type, the constant becomes typed.  Constants must be
declared as their correct type, or else the program will return an error.

Declare constants without a type unless you absolutely need them;  declaring
a type makes you lose the flexibility of being able to mix types in an 
operation.

#### Rules for constant expressions in Go

- Comparison between two untyped constants results in an untyped boolean constant ("true"/"false").
- Operands of the same type result in the same type. "11/2" results in "5" rather than "5.5" because it's truncated into an integer.
- If they aren't the same type, the result is the broader of the two according to this logic: integer<rune<floating-point<complex.

### Variables

A variable is the name given to a particular storage area used by the program.  Names
can use letters, digits, and underscores.  They have a specific size, memory layout,
range of values for the memory layout, and possible operations.  The definition tells
the compiler how much storage to create for the variable and where to put it.

*Lvalues* are used to refer to a memory location and can appear on either the left or
right side of an assignment. 

*RValues* refer to data values that are stored in memory.

Variables are lvlaues and appear on the left side of the assignment.

When you declare a variable you'll use the keyward var.

### Operators

Operators are built in symbols that perform a variety of mathematical or logical operations.

Three types of Operators:

- bitwise
- arithmetic
- logical

#### Bitwise Operators

Operate on variables of the same bit length (that are integers).

- AND: &
- OR: |
- XOR and Complement: ^
- CLEAR: &^

#### Arithmetic Operators

- "+" (addition)
- "*" (multiplication)
- "/" (division)
- "%" (percentage)

#### Logical Operators

- Equal: "="
- Not Equal: "!="
- Less Than: "<"
- Less than or equal to: "<="
- Greater than: ">"
- Greater than or equal to ">="

### Strings

Strings are differeent from the way they work in C++, Python, Java, etc.

**Strings are immutable, or read-only.**

If you try to change a string, you'll get an error.  The characters represent 
bytes that are UTF-8 encoded.  They are delineated by double quotes.

Example:

```Golang
package main

import "fmt"

func main() {
	var my_words string
	
	my_words = "Hello World!"
	
	fmt.Println("String: ", my_words)
}
```

Note: You can use backticks when writing in HTML, when writing a message that 
takes multiple lines, and when writing regular expressions.

### Times and Dates

In Go, times include a location that determines the date and time associated
with that location.  If it's not specified, the time defaults to UTC.

If you want a timestamp, Time.Now, the signature is

> func Now() time

For the date and time, Time.Date, the signature is in the format yyyy-mm-dd hh-mm-ss + nanoseconds:

> func Date (year int, month, day, hour, min, sec, nsec int, loc *Location)

In Go, the duration is elapsed in nanoseconds between two instants written in int64nanosecond count.

1 second is equal to 1,000,000,000 nanoseconds.

The function for duration is:

> func Since(t Time) duration.

If you want to know how long until time t, the function is:

> func Until(t Time) duration


## Control Structures

Control structures allow the program to determine where to go next or which 
path to follow.  They analyze the variables and figure out where to go 
based on the given parameters.

Types of control structures

- Conditional (selection)
- Loops (iteration)

### If-Else

```Golang
If condition {
	// code that will be executed if the condition is true
}
```

Example

```Golang
package main

import (
	"fmt"
)

func main () {
	var grade = "A"
	
	x := true
	
	if x {
		fmt.Println(grade)
    }
}
```

```Golang
package main

func main() {
	var condition = true
	
	if condition {
		// code that will be executed if the condition is true
	} else {
		// code that will be executed if the condition is false
	}
}
```

Example

```Golang
package main

import (
	"fmt"
)

func main() {
	x := 75
	
	if x >= 60 {
		fmt.Println("Pass")
    } else {
		fmt.Println("Fail")
    }
}
```

```Golang
package main

func main() {
	var condition_1 = true
	var condition_2 = false
	
	if condition_1 {
		// code that will be executed if the condition_1 is true
	} else if condition_2 {
        // code that will be executed if condition_2 is true
	} else {
		// code that will be executed if both conditions are false
	}
}
```

### Switch-Case

In Go the control switch runs the first case where the condition is true and
ignores the remainder.  Switch cases don't need to be constants, and the 
values don't have to be integers.  You can use commas to separate expressions
int he case if there is more than one.

Example

```Golang
package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
```

### For

A *for loop* is the only looping construct in Go.  It's used to repeat statements
until the condition is met.  If there was no condition to be met.

Example

```Golang
package main

import (
	"fmt"
)

func main() {
	x := 1
	
	for x <= 10 {
		fmt.Println(x)
		
		x += 1
    }
}
```

```Golang
package main

import (
	"fmt"
)

func main() {
	x := 1
	
	for x := 1 ; x <= 10 ; x++ {
		fmt.Println(x)
    }
}
```

Example: Adding numbers in a slice

```Golang
package main

import (
	"fmt"
)

func main() {
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
    }
	
	fmt.Println("Sum: ", sum)
}
```

Note: break and continue can be utilized with conditionals to alter loop flow

## Functions

Function names that start with lowercase will only be used within the same 
package.  A package is a group of source files that are compiled together.
Names of functions can't start with a digit and can't contain spaces.  If
the name has multiple words in it, each word after the first should be 
capitalized, such as simpleFunction.

### Basics of Functions

Example: A function that doesn't have any parameters and doesn't return any 
values.

```Golang
func SimpleFunction() {
	fmt.Println("Hello World!")
}

func main() {
	SimpleFunction()
}
```

Example: A function that passses arguments

```Golang
package main

import (
	"fmt"
)

func add(x int, y int) {
	total := x + y
	
	fmt.Println(total)
}

func main() {
	add(10, 20)
}
```

### Return Values and Parameters

If you declare a return value, then you'll also need a return statement.  The
*return value* is what retrains the function's output after running the 
instructions.  In Go there's a variety of values.  You can decide whether you
want to name it or not.

Example: Naked return

```Golang
package main

import (
	"fmt"
)

func greeting(name string) (message string) {
	message = "Hello, " + name // variable has to be the same name in the header
	
	return // since message has een specified no need to declare here
}

func main() {
	msg := greeting("world")
	
	fmt.Println(msg)
}
```

## Examples

- [Hello](./hello/README.md)

# Environment Setup

> export GOROOT=/usr/local/go

> export PATH=$PATH:$GOROOT/bin


Initializing the root module

> go mod init golab.com/m/v2

Adding module requirements and sums

> go mod tidy

Executing tests

> cd math
> go test

Executing tests with coverage

> cd math
> go test -v --cover

Installing mockgen

> go install go.uber.org/mock/mockgen@latest

> go get github.com/uber/mock/gomock
> go get go.uber.org/mock/mockgen/model

Using mockgen to generate mocks (from the root)

> mockgen -destination=mocks/mockRunner.go -package=mocks golab.com/m/v2/IUser IUserInterface

## References

- [Playground](https://go.dev/play/)
- [Home](https://golang.org)
- [User Manual](https://go.dev/doc)
- [Standard Library](https://pkg.go.dev/std)
- [Effective Go](https://go.dev/doc/effective_go)
- [Recorded Talks](https://go.dev/talks/)
- [Frequently Asked Questions](https://go.dev/doc/faq)
- [Testing with GoMock](https://gist.github.com/thiagozs/4276432d12c2e5b152ea15b3f8b0012e)
- [Using Go Modules](https://go.dev/blog/using-go-modules)
