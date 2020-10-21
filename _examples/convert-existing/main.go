package main

import (
	"os"

	"github.com/arjit95/cobi"
	"github.com/arjit95/cobi/editor"
	"github.com/spf13/cobra"
)

var (
	cmd         *cobi.Command
	interactive bool
)

// This example demonstrates how to add interactive prompt support to an existing app
func init() {
	cmd1 := &cobra.Command{
		Use:   "demo",
		Short: "Simple demo command",
	}

	pods := []string{"pod-11", "pod-12", "pod-2"}

	cmd2 := &cobra.Command{
		Use:   "pods",
		Short: "List all pods",
		Run: func(cmd *cobra.Command, args []string) {
			for _, pod := range pods {
				cmd.Printf("%s\n", pod)
			}
		},
	}

	cmd3 := &cobra.Command{
		Use:       "get",
		Short:     "Get pod info",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: pods,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("%s info\n", args[0])
		},
	}

	cmd2.AddCommand(cmd3)
	cmd1.AddCommand(cmd2)

	// Add an interactive flag and wrap your command with cobi
	cmd1.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run shell in interactive mode")
	cmd = cobi.NewCommand(editor.NewEditor(), cmd1)
}

func main() {
	err := cmd.ParseFlags(os.Args)
	if err != nil {
		cmd.PrintErrf("%s\n", err.Error())
		cmd.Usage()
		os.Exit(0)
	}

	if interactive {
		cmd.Flag("interactive").Hidden = true
		cmd.ExecuteInteractive()
	} else {
		cmd.Execute()
	}
}
