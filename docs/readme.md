# Go

## Ssetting Up Your Environment

### Setting GOROOT to Define the Go Binary Location

The operating system needs to know how to find the Go installation.  In most instances, if you've installed Go in the 
default path, such as /usr/local/go on a Unix system, you don't have to take any action.  However, in the event that
you've chosen to install Go in a nonstandard path or are installing Go on Windows, you'll need to tell the operating
system where to find the Go binary.

You can do this from your command line by setting the reserved ```GOROOT``` environment variable to the location of your
binary.  


### Setting GOPATH to Define the Go Workspace

Unlike setting your ```GOROOT``` environment variable, which is optional, you must set your ```GOPATH``` environment
to instruct the Go toolset where your source code, third party libraries, and compiled programs will exist.

The ```GOPATH``` environment variable is a list of paths that point to your Go workspace.  The Go workspace is a 
directory that contains your Go source code, third party libraries, and compiled Go programs.  The ```GOPATH``` 
contains three subdirectories within: bin, pkg, and src.

The bin directory will contain your compiled and installed Go executable binaries.  Binaries that are built and 
installed will be automatically placed in this location.

The pkg directory stores various package objects including thrid-party Go dependencies that your code might rely on.

The src directory will contain all the source code you'll write.

## SOLID

The SOLID principles are essentially a set of rules for helping you write clean
and maintainable object-oriented code.

* Single responsibility
* Open/closed
* Liskov substitution
* Interface segregation
* Dependency inversion

Go is an object-based programming language.

## The Single Responsibility Principle (SRP)

Object implementations should focus on doing one thing well, and in an efficient
way.

## Open/closed

## Liskov substitution

## Interface segregation

## Dependency inversion

