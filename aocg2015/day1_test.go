package aocgtwft

import (
    "fmt"
    "testing"
)

func TestDayOne(t *testing.T) {
    type partOneTestCase struct {
        input string
        expected int
    }

    type partTwoTestCase struct {
        input string
        expected int
    }

    partOneTests := []partOneTestCase{
        {"(())", 0},
        {"))(((((", 3},
    }

    partTwoTests := []partTwoTestCase{
        {")", 0},
        {"()())((", 5},
    }

    passCount := 0
    failCount := 0
    
    fmt.Printf("\n---- Day One :: Part One ----\n\n")

    for _, test := range partOneTests {
        floorOutput := partOne(test.input)
        if floorOutput != test.expected {
            failCount++
            t.Errorf(`-------------------------------
Input: %v
Expected: %v
Actual: %v
Fail
`, test.input, test.expected, floorOutput)
        } else {
            passCount++
            fmt.Printf(`-----------------------------
Input: %v
Expected: %v
Actual: %v
Pass
`, test.input, test.expected, floorOutput)
        }
    }

    fmt.Printf("-----------------------------\n")
    fmt.Printf("%d passed, %d failed\n", passCount, failCount)

    passCount = 0
    failCount = 0

    fmt.Printf("\n----- Day One :: Part Two ------\n\n")
    
    for _, test := range partTwoTests {
        basementOutput := partTwo(test.input)
        if basementOutput != test.expected {
            failCount++
            t.Errorf(`------------------------------
Input: %v
Expected: %v
Actual: %v
Fail
`, test.input, test.expected, basementOutput)
        } else {
            passCount++
            fmt.Printf(`----------------------------
Input: %v
Expected: %v
Actual: %v
Pass
`, test.input, test.expected, basementOutput)
        }
    }

    fmt.Printf("---------------------------\n")
    fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
