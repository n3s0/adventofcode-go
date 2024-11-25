package aocg2015

import (
    "fmt"
    "testing"
)

func TestDayTwo(t *testing.T) {
    type partOneTestCase struct {
        input string
        expected int
    }
    
    type partTwoTestCase struct {
        input string
        expected int
    }

    partOneTest := []partOneTestCase{
        {`2x3x4`, 58},
        {`1x1x10`, 43},
    }

    partTwoTest := []partTwoTestCase{
        {`2x3x4`, 34},
        {`1x1x10`, 14},
    }

    passCount := 0
    failCount := 0
    
    fmt.Printf("\n---- Day Two :: Part One ----\n")

    for _, test := range partOneTest {
        squareFootageOutput := dayTwoPartOne(test.input)
        if squareFootageOutput != test.expected {
            failCount++
            t.Errorf(`-------------------------------
Inputs: %v
Expected: %v
Actual: %v
Fail
`, test.input, test.expected, squareFootageOutput)
        } else {
            passCount++
            fmt.Printf(`-----------------------------
Inputs: %v
Expecting: %v
Actual: %v
Pass
`, test.input, test.expected, squareFootageOutput)
        }
    }

    fmt.Println("-----------------------------")
    fmt.Printf("%d passed, %d failed\n", passCount, failCount)
    fmt.Println("-----------------------------")

    passCount = 0
    failCount = 0
    
    fmt.Printf("\n---- Day Two :: Part Two ----\n")

    for _, test := range partTwoTest {
        squareFootageOutput := dayTwoPartTwo(test.input)
        if squareFootageOutput != test.expected {
            failCount++
            t.Errorf(`-------------------------------
Inputs: %v
Expected: %v
Actual: %v
Fail
`, test.input, test.expected, squareFootageOutput)
        } else {
            passCount++
            fmt.Printf(`-----------------------------
Inputs: %v
Expecting: %v
Actual: %v
Pass
`, test.input, test.expected, squareFootageOutput)
        }
    }

    fmt.Println("-----------------------------")
    fmt.Printf("%d passed, %d failed\n", passCount, failCount)
    fmt.Println("-----------------------------")
}
