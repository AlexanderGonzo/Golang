Modules
- in GoLang is how we can track dependencies that we may have.
- Go's way of package management.
- go.mod file lists all dependencies, declare version of go.

Go Random Notes
- go.sum - details of dependencies (modules).
- go run main.go doesn't not build any artifacts.
- go build main.go will build a binary as the name of the file.
- go build will pull everything the go file dependencies on and name it as the project name.
- go install saves binary in GO PATH binary.
- If we want to make a library that can be imported into GO projects.
    - make a file that matches the name of the directory (package name same as directory)
    - instead of package main we do package filename
- No scope separation between GO files.
- Anything with _test is considered a test for something.

Function Notes
- When a Function is written with a capital letter the capital letter will get exported automatically.

Test
- Can write examples that can be used as tests that people can use
- examples can also be ran dynamically on the playground that is documenting your code

Packages
- Basically a github repo or a collection of files that we import into our application

Go Overview
Go aims to reduce typing and complexity using a minimal amount of keywords, so you will code less than in other languages like Java. Keywords can be parsed without a symbol table, as the grammar is LALR(1). Go acts like a hybrid, imperative language, but it is built with concurrency in mind. Here are some of the unique features of Go:

- No function or operator overloading
- No implicit conversions to avoid bugs
- No classes or type inheritance
- No variant types
- No dynamic code loading or dynamic libraries
- No assertions or immutable variables

Go is a compiled language there is no need for an interrupter. Go, C, and Rust are all languages where the code is first converted to machine code by the compiler before it's executed. Code is ran from top to bottom.

Go Basics Topics
- Filenames, keywords, identifiers
- Operators, types, functions, and constants
- Pointers, structures, methods
- Maps, arrays, slices
- Go CLI
- Interface
- Error handling
- Goroutine, Channel, Buffer
- Panic, Defer, Error, Recover
- Go design patterns
- Variables
    - You need to declare a variable with a unique name "identifier" before you can use it.
    - Naming Rules. Also apply to function and package.
        - Names always start with a letter or an uppercase letter
        - Upper case letters are exported
        - Underscore characters
        - unicode letters
        - Names cannot start with numbers or punctuation  
    - Every Variable and value has a type. Go is a strongly-typed language. (static type)
    - Static Types: cannot change type after declaration 
    - Blank identifiers discard values.
- Naming Convention 
    - Naming is important because readability = maintainability 
    - use the first few letters of the word 
    ```go
    var fv string // flag value 
    ```
    - use few letters in smaller scopes 
    ```go
    var bytesRead int // number of bytes read XXX DO NOT DO XX
    var n int // number of bytes read !! DO THIS !!
    ```
    - use the complete words in larger scopes 
    ```go
    package file 
    var fileClosed bool 
    ```
    - Use mixedCaps like this
    ```go
    type PlayerScore struct 
    ```
    - Use all capitals for acronyms
     ```go
    var localAPI string // GOOD
    var localApi string // BAD 
    ```
    - Do not Stutter 
     ```go
    player.PlayerScore // BAD
    player.Score // GOOD  
    ```
    - Do not use under_scores or LIKE_THIS
    ```go
    const MAX_TIME int // BAD
    const MaxTime int // GOOD
    const N int // GOOD 
    ```
    - Not Idiomatic "unnecessarily verbose"
    ```go
        func Read(buffer *Buffer, inBuffer []byte) (size int, err error) {
            if buffer.empty() {
                buffer.Reset() 
            }
            size = copy( 
        inBuffer, 
        buffer.buffer[buffer.offset:])
            buffer.offset += size
            return size, nil
        }
    ```
    - Idiomatic "concise and idiomatic" 
       ```go
        func Read(b *Buffer, p []byte) (n int, err error) {
            if b.empty() {
                b.Reset() 
            }
            n = copy(p, b.buf[b.off:])
            b.off += n
            return n, nil
        }
        ```
    - Common Abbreviations used in Go
        ```go
        var s string // string
        var i int // index
        var msg string // message
        var v string // value
        var val string // value
        var num int // number
        var fv string // flag value
        var err error // error value
        var args []string // arguments
        var seen bool // has seen?
        var parsed bool // parsing ok?
        var buf []byte // buffer
        var off int // offset
        var op int // operation
        var opRead int // read operation
        var l int // length
        var n int // number or number of
        var m int // another number
        var c int // capacity
        var c int // character
        var a int // array
        var r rune // rune
        var sep string // separator
        var src int // source
        var dst int // destination
        var b byte // byte
        var b []byte // buffer
        var buf []byte //buffer
        var w io.Writer //writer 
        var r io.Reader // reader 
        var pos int //position 
        ...list goes on and on..
        ```
- Short VS Normal Declaration
    - Normal Declaration
        - If you do not know the initial value
        - When you need a package scoped variable
        - When you want to group variables together for greater readability
    - Short Declaration 
        - Most used and preferred declaration in Go
        - If you know the initial value
        - To keep the code concise and easy to read
        - For re-declaration  
        - Can be used inside if and switch statements that are scoped
    - Note
        - Beware of Shadowing 
- Type Conversion 
    - type(value)
- Introduction to Slices
    - A slice can store multiple values 
    ```go
    var Args []string
    ```
    - Here Arg's type is a "slice of string". Which means args can store a series of string values inside
    - Args can store multiple string values
    - but the slice itself is a single value it just points to the other values it stores
    - Each value inside a slice is an unnamed variable. 
    - How do we access the value w/out the name? We can use index expressions 
    ```go
    go run main.go hi yo
    Args[0] // Stores path to the program
    Args[1] // Stores first argument "hi"
    Args[2] // Stores second argument "yo"
    ```
- Raw String Literals
    - characters are wrapped with `hi` and can contain multiple lines of text
    - Unlike string literals the first value in a raw string literal is not interpreted.
    - raw string literals is an unprocessed string data.
    ```go
    var hi string = `hi
                    there!`
    ```
- IOTA
    - iota is a built-on constant generator which generates ever increasing numbers 
    - iota starts at 0, unless expressed otherwise
    - You can use an expression with iota. So, the other constants will repeat the express 
    ```go
        const(
            monday = 0
            tuesday = 1
            wednesday = 2
            ..
            sunday = 6
        )
        fmt.Println(monday, ..., sunday) // 0 1 2 3 4 5 6 

        //IOTA
        const (
            monday = iota
            tuesday
            wednesday
            sunday 
        )
        fmt.Println(monday, ... , sunday) // 0 1 2 3 4 5 6 

        const (
            monday = iota + 1
            tuesday
            wednesday
            sunday 
        )
        fmt.Println(monday, ... , sunday) // 1 2 3 4 5 6 7 
    ```
    - blank identifer 
    ```go
        const (
            EST = -5
            MST = -7
            PST = -8
        )
        fmt.Println(EST, MST, PST) // -5 -7 -8

        const (
            EST = -(5 + iota)
            MST  
            PST  
        )
        fmt.Println(EST, MST, PST) // -5 -6 -7

        const (
            EST = -(5 + iota)
            _
            MST  
            PST  
        )
        fmt.Println(EST, MST, PST) // -5 -7 -8
    ```


Go Advanced Topics
- Go dependency management tools
- Semantic versioning
- Scripts and repositories
- Go libraries
    - Go Kit, GORM, Gen, and CLI.
- SQL fundamentals
- GIT
- Basic authentication
- HTTP/HTTPS
- Web frameworks and routers
    - Frameworks: Echo, Beego, Gin, Revel, and Chi, with Echo being the most important for Go.
- Relational databases (PostgreSQL)
- Log Frameworks: Zap

Go Testing
- Unit Testing
    - Built-in testing packages.
- Integration Testing
- Behavior Testing
- E2E testing
- Testing Frameworks
    - Ginkgo and GoCheck
    - Ginkgo can also be used for behavior and integration testing.

Go Patterns
- Structural
- Creational
    - Builder, Factory, and Singleton
- Behavioral 
    - Iterator, Observer, and Command
- Concurrency
    - Adapter, Bridge, and Decorator
- Stability

Common Go Interview Questions
- What is a goroutine? How do you stop it?
- How do you check a variable’s type at runtime?
- How do you format a string without printing?
- How do you concatenate strings in Go?
- What is Go 2?
- How do you initialize a struct in Go?

Resources
https://www.openmymind.net/assets/go/go.pdf
https://github.com/golang/go/wiki/Learn
https://betterprogramming.pub/top-5-common-golang-coding-mistakes-the-ugly-sides-of-a-great-programming-language-e0b64915707
