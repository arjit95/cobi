package main

import (
	"os"

	"github.com/arjit95/cobi"
	"github.com/spf13/cobra"
)

var (
	cmd         *cobi.Command
	interactive bool
)

func init() {
	cmd = &cobi.Command{
		Command: &cobra.Command{
			Use:   "demo",
			Short: "Simple demo command",
		},
	}

	pods := []string{"pod-11", "pod-12", "pod-2"}

	cmd2 := &cobi.Command{
		Command: &cobra.Command{
			Use:   "pods",
			Short: "List all pods",
			Run: func(cmd *cobra.Command, args []string) {
				for _, pod := range pods {
					cmd.Printf("%s\n", pod)
				}
			},
		},
	}

	cmd3 := &cobi.Command{
		Command: &cobra.Command{
			Use:       "get",
			Short:     "Get pod info",
			Args:      cobra.ExactValidArgs(1),
			ValidArgs: pods,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Printf("%s info\n", args[0])
			},
		},
	}

	cmd2.AddCommand(cmd3)
	cmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run shell in interactive mode")
	cmd.AddCommand(cmd2)
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
		cmd.BuildInteractiveSession(true)
	} else {
		cmd.Execute()
	}
}
