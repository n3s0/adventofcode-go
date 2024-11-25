package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of adventofcode-go",
  Long: `Print the version number of adventofcode-go. Prints the 
            current year, day, and part that is currently being 
            worked on. (e.g. 15.2.0 means 2015 day 2 is being worked 
            on. 2015.3.2 means 2015 day 3 part 2 is being worked on.`,
  Run: func(cmd *cobra.Command, args []string) {
        const aocgVersion = "15.3.1"
        fmt.Printf("adventofcode-go v%s", aocgVersion)
  },
}
