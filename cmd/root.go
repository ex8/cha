package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ex8/cha/http"
	"github.com/spf13/cobra"
)

var client = http.New()

var rootCmd = &cobra.Command{Use: "cha"}

var commands = []*cobra.Command{
	donationsCmd,
	nonprofitsCmd,
	climateCmd,
}

func Execute() {
	rootCmd.AddCommand(commands...)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func prettyPrint(o interface{}) string {
	s, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(s)
}
