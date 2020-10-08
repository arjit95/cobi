package main

import (
	"fmt"

	"github.com/arjit95/cobi"
	"github.com/spf13/cobra"
)

var (
	cmd *cobi.Command
)

func init() {
	cmd = cobi.NewCommand(&cobra.Command{
		Use:   "logger",
		Short: "Simple demo command",
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Log messasge to logger",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			fmt.Fprint(cmd.Editor.Logger.Info, args[0]+"\n")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "error",
		Short: "Log messasge to logger",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			fmt.Fprint(cmd.Editor.Logger.Error, args[0]+"\n")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "warn",
		Short: "Log messasge to logger",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			fmt.Fprint(cmd.Editor.Logger.Warn, args[0]+"\n")
		},
	})
}

func main() {
	cmd.ExecuteInteractive()
}
