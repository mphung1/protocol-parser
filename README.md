# Running

From this directory, there are two options to run the parser:

Specifying the list of input values to the executable:

Interpreted
```
go run main.go 1 2 3 4 5
```

Specifying the relative path to an examples file:

Interpreted
```
go run main.go -f examples/example-1.txt
```

Input values in a test file are delimited by whitespace.

# Testing

The project will be graded using unit tests. After each phase, the tests used to grade the phase will be
added to your repository. You should ensure these tests continue to pass; you will be penalized if you break
previous tests in a subsequent phase.

For Go, all grading tests will be in `parser_test.go`. Do not change this file; it will
be overwritten each phase with the new grading tests.

To run the grading tests:
```go
go test -v ./...
```

# Notes

* Do not change any code in the `main.go` file.
* You have free rein on the rest of the application, you may create other methods, structs, or files as necessary. For simplicity, please keep new files in the same directory.
* Do not delete the existing `example-1.txt` file.
* Other test files may be added to the `examples` directory and committed to your repository.
