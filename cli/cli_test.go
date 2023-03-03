package cli

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestCLI_Add(_ *testing.T) {
	cli := NewCLI("test")

	root := &cobra.Command{
		Use: "cli_test",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("pre run")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("post run")
		},
	}
	cli.SetRoot(root)

	cli.Add("foo.bar", &cobra.Command{
		Use: "bar",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cli_test.foo.bar")
			return nil
		},
	})

	cli.Add("foo.baz", &cobra.Command{
		Use: "baz",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("cli_test.foo.baz")
			return nil
		},
	})

	fmt.Println("asdf")
}
