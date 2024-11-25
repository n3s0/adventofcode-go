package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "adventofcode-go",
  Short: "Advent of Code Go shows my progress in Advent of Code challenges.",
  Long: `Advent of Code Go shows my progress and PoC code as I work
            throught the Advent of Code challenges. To review the PoC
            please go to https://github.com/n3s0/adventofcode-go`,
  Run: func(cmd *cobra.Command, args []string) {
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
