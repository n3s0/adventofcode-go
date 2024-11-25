package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/n3s0/adventofcode-go/aocg2015"
)

var dayone bool
var daytwo bool

func init() {
    rootCmd.AddCommand(aocg2015Cmd)

    aocg2015Cmd.Flags().BoolVarP(&dayone, "dayone", "1", false, "day one solution(s)")
    aocg2015Cmd.Flags().BoolVarP(&daytwo, "daytwo", "2", false, "day two solution(s)")
}

var aocg2015Cmd = &cobra.Command{
    Use:   "2015",
    Short: "Provides the solutions for Advent of Code 2015",
    Long:  `Provide solutions for Advent of Code 2015.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Advent of Code Go :: 2015")

        if len(args) < 1 {
            fmt.Println("No input provided...")
        }

    },
}
