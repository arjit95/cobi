package main

import (
	"fmt"
	"os"

	"github.com/arjit95/cobi"
	"github.com/arjit95/cobi/editor"
	"github.com/spf13/cobra"
)

var (
	cmd         *cobi.Command
	interactive bool
)

func init() {
	cmd = cobi.NewCommand(&cobra.Command{
		Use:   "demo",
		Short: "Simple demo command",
	}, editor.NewEditor())

	pods := []string{"pod-11", "pod-12", "pod-2"}

	cmd.AddCommand(&cobra.Command{
		Use:   "pods",
		Short: "List all pods",
		Run: func(cmd *cobra.Command, args []string) {
			for _, pod := range pods {
				cmd.Printf("%s\n", pod)
			}
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:       "get",
		Short:     "Get pod info",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: pods,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("%s info\n", args[0])
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "log",
		Short: "Log messasge to logger",
		Args:  cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			if cmd.InteractiveMode() {
				fmt.Fprint(cmd.Editor.Logger.Info, args[0]+"\n")
			} else {
				fmt.Printf("[Info] %s\n", args[0])
			}
		},
	})

	cmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run shell in interactive mode")
}

func main() {
	err := cmd.ParseFlags(os.Args)
	if err != nil {
		cmd.PrintErrf("%s\n", err.Error())
		cmd.Usage()
		os.Exit(0)
	}

	if interactive {
		cmd.Editor.SetUpperPaneTitle("Cobi Demo")
		cmd.Flag("interactive").Hidden = true
		cmd.ExecuteInteractive()
	} else {
		cmd.Execute()
	}
}
