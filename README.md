# Go Panic and Recover Example

This repository demonstrates a common error related to Go's `panic` and `recover` mechanisms.  The `bug.go` file contains code that panics.  The `bugSolution.go` file shows the correct way to handle panics and prevent unexpected behavior.

## Problem

Many developers assume that `recover` will allow code execution to continue after a `panic`.  This is incorrect.  Code following a `panic` will not execute, even if a `recover` function is in place.  This example highlights that issue.