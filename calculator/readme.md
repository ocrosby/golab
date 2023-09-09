# Calculator

## Installing Ginkgo

### Step 1. Setup

Add the following to .zshrc

```shell
export GOROOT=~/go
export GOBIN=${GOROOT}/bin

export PATH=${GOBIN}:${PATH}
```

### Step 2. Installing the Ginkgo and Gomega libraries
> go get github.com/onsi/ginkgo/v2/ginkgo

> go get github.com/onsi/gomega/...

### Step 3. Installing the Ginkgo program
> go install github.com/onsi/ginkgo/v2/ginkgo


### bootstrapping the test suite

change into the calculator package directory

> ginkgo bootstrap

This results in a file named <package>_suit_test.go
to be created which is sort of a starting point for 
testing the package using the ginkgo tool.


### Generating a new test fixture for our target module

> ginkgo generate calculator

this generated a test file named calculator_test.go with the describe boilerplate

