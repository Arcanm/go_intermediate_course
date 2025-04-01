// EXECUTE NEXT LINE TO RUN TEST OF ALL PROJECT //
// go test

// --------------------------- //
// ------ CODE COVERAGE ------ //
// --------------------------- //
// TO OBTAIN CODE COVERAGE OF THE PROJECT YOU CAN RUN:
// go test -cover -> ONLY RETURNS THE PERCENTAGE OF CODE COVERAGE

// TO OBTAIN MORE DETAILS ABOUT THE MISSING COVERAGE
// go test -coverprofile=coverage.out -> NOT LEGIBLE FILE

// OPEN THE coverage.out IN CONSOLE OR WEB BROWSER
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out

// ----------------------------- //
// --------- PROFILING --------- //
// ----------------------------- //
// TO OBTAIN PROFILING OF THE PROJECT YOU CAN RUN:
// go test -cpuprofile=cpu.out

// TO ANALIZE THE RESULTS IN CONSOLE YOU CAN RUN THE COMMAND:
// go tool pprof cpu.out

// INSIDE OF PPROF //
// top -> Show you the times of execution of the testing running
// list [funcName] -> Obtain details of execution, for example
// list Fibonacci
// web -> Open web diagram with details

package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)

	// NOT COMMON TESTING TYPE
	if total != 10 {
		t.Errorf("Sum func was incorrect, got %d expected %d", total, 10)
	}

	// COMMON TESTING PATTERN
	tables := []struct {
		a             int
		b             int
		expectedTotal int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.expectedTotal {
			t.Errorf("Sum func was incorrect, got %d expected %d", total, item.expectedTotal)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a           int
		b           int
		expectedMax int
	}{
		{5, 2, 5},
		{3, 2, 3},
		{2, 3, 3},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.expectedMax {
			t.Errorf("GetMax func was incorrect, got %d expected %d", max, item.expectedMax)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := []struct {
		a              int
		expectedResult int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.a)

		if item.expectedResult != fib {
			t.Errorf("Fibonacco was incorrect, got %d expected %d", fib, item.expectedResult)
		}
	}
}
