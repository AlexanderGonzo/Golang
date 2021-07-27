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
- 

Test
- Can write examples that can be used as tests that people can use
- examples can also be ran dynamically on the playground that is documenting your code

Packages
- Basically a github repo or a collection of files that we import into our application

Go aims to reduce typing and complexity using a minimal amount of keywords, so you will code less than in other languages like Java. Keywords can be parsed without a symbol table, as the grammar is LALR(1). Go acts like a hybrid, imperative language, but it is built with concurrency in mind. Here are some of the unique features of Go:

- No function or operator overloading
- No implicit conversions to avoid bugs
- No classes or type inheritance
- No variant types
- No dynamic code loading or dynamic libraries
- No assertions or immutable variables

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