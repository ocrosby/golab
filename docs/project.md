# Projects

What does it mean to be 'go gettable'?

## Code Organization

Go programs are organized into packages. A package is a collection of source files in the same
directory that are compiled together.  Functions, types, variables, and constants defined in one
source file are visible ot all other source files in the same package.

A repository contains on e or more modules.  A module is a collection of related Go packages that
are released together. A Go repository typically contains one module, located at the root of the
repository.  A file named go.mod there declares the module path: the import path prefix for all
packages within the module.  The module contains the packages in the directory containing its go.mod
file as well as subdirectories of that directory, up to the next subdirectory containing another go.mod
file (if any).

Note that you don't need to publish your code to a remote repository before you can build it.  A module
can be defined locally without belonging to a repository.  However, it's a good habit to organize your 
code as if you will publish it someday.

Each module's path not only serves as an import path prefix for it's packages, but also indicates 
where the go command would consult the repository indicated by https://golang.org/x/tools.

An import path is a string used to import a package.  A package's import path is its module path
joined with its subdirectory within the module.  For example the module github.com/google/go-cmp
contains a package in the directory cmp/.  That package's import path is github.com/google/go-cmp/cmp.

There are a number of project layout patterns emgerging in the Go echosystem.  The two most common patterns 
are:

* cmd directories
* pkg directories

The cmd layout pattern is very useful when you need to have more than one application binary. 
Each binary gets a subdirectory (e.g. your_project/cmd/your_app).  This pattern also helps you keep your
project/package 'go gettable'.  That means you can use the go get command to fetch (and install) your
project, it's applications and its libraries (e.g. go get github.com/your_github_username/your_project/cmd/appxg).
You don't have to separate the applciation files.  You'll be able to build each application with the right
set of go build flags, but go get will no longer work because it will not know which application code 
to build.  The official Go tools is one example of the cmd layout pattern.

Some projects that use the cmd pattern:

* Kubernetes
* Docker
* Prometheus
* Influxdb

The pkg layout pattern is also pretty popular.  For new Go developers it's one
of the most confusing package structure concepts because Go workspaces have a directory
with the same name and that directory has a different purpose (it's used to store object files for the
packages the Go compiler builds).  The pkg directory is where you put your public libraries.  They
can be used internally by your application.  They can also be used by external projects.  This
is an informal contract between you and the other external users of your code. Other projects will
import these libraries expecting them to work, so think twice before you put something here.

### Importing packages from remote modules

An import path can describe how to obtain the package source code using a revision control system
such as Git or Mercurial.  The go tool uses this property to automatically fetch packages from remote
repositories.

## Creating a Go Project

## Project Layout

## Go Workspaces



## References

* [How to Write Go Code](https://golang.org/doc/code.html)
* [Go Project Layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)
* [Tips to Create a Proper Go Project Layout](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2)
* [An Analysis of the Top 1000 Go Repositories](http://blog.sgmansfield.com/2016/01/an-analysis-of-the-top-1000-go-repositories/)