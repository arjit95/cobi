package main

import (
	"fmt"
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
		RootCmd: &cobra.Command{
			Use:   "demo",
			Short: "Simple demo command",
		},
	}

	pods := []string{"pod-1", "pod-2"}

	cmd2 := &cobi.Command{
		RootCmd: &cobra.Command{
			Use:   "pods",
			Short: "List all pods",
			Run: func(cmd *cobra.Command, args []string) {
				for _, pod := range pods {
					fmt.Printf("%s\n", pod)
				}
			},
		},
	}

	cmd3 := &cobi.Command{
		RootCmd: &cobra.Command{
			Use:       "get",
			Short:     "Get pod info",
			Args:      cobra.ExactValidArgs(1),
			ValidArgs: pods,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("%s info\n", args[0])
			},
		},
	}

	cmd2.AddCommand(cmd3)
	cmd.RootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run shell in interactive mode")
	cmd.AddCommand(cmd2)
}

func main() {
	err := cmd.RootCmd.ParseFlags(os.Args)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		cmd.RootCmd.Usage()
		os.Exit(0)
	}

	if interactive {
		cmd.RootCmd.Flag("interactive").Hidden = true
		cmd.InitDefaultExitCmd()
		cmd.RunInteractive()
	} else {
		cmd.Execute()
	}
}
