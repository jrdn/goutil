package cli

import (
	"strings"

	"github.com/j13g/goutil/types"
	"github.com/samber/mo"
	"github.com/spf13/cobra"
)

type cNode struct {
	name string
	cmd  mo.Option[*cobra.Command]
}

func newCLINode(name string, cmd *cobra.Command) *cNode {
	var o mo.Option[*cobra.Command]
	if cmd != nil {
		o = mo.Some(cmd)
	} else {
		o = mo.None[*cobra.Command]()
	}
	return &cNode{name: name, cmd: o}
}

type CLI struct {
	appName  string
	commands types.Tree[*cNode]
}

func NewCLI(appName string) *CLI {
	cli := &CLI{
		appName:  appName,
		commands: types.NewTree[*cNode](),
	}
	cli.SetRoot(&cobra.Command{Use: appName})
	return cli
}

func (c *CLI) SetRoot(cmd *cobra.Command) *CLI {
	c.commands.Set(newCLINode(c.appName, cmd))
	return c
}

func (c *CLI) Add(path string, command *cobra.Command) *CLI {
	parts := strings.Split(path, ".")

	node := c.commands

	// iterate over all but the last part
parts:
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		for _, child := range node.Children() {
			if child.Get().name == part {
				node = child
				continue parts
			}
		}

		node = node.Add(newCLINode(part, nil))
	}

	name := parts[len(parts)-1]
	node.Add(newCLINode(name, command))
	return c
}

func (c *CLI) Run() error {
	rootTreeNode := c.commands
	rootCLINode := rootTreeNode.Get()

	if rootCLINode == nil {
		cmd := &cobra.Command{Use: c.appName}
		rootCLINode = newCLINode(c.appName, cmd)
		rootTreeNode.Set(rootCLINode)
	}

	if rootCLINode.cmd.IsAbsent() {
		rootCLINode.cmd = mo.Some(&cobra.Command{
			Use: c.appName,
		})
	}

	c.walk(rootTreeNode)

	// run the CLI
	cmd, _ := rootCLINode.cmd.Get()
	return cmd.Execute()
}

// walk tree to set children on cobra commands and create placeholders for sparse nodes in the tree
func (c *CLI) walk(treeNode types.Tree[*cNode]) {
	node := treeNode.Get()

	cmd, set := node.cmd.Get()
	if !set {
		cmd = &cobra.Command{Use: node.name}
	}

	for _, childTreeNode := range treeNode.Children() {
		childCLINode := childTreeNode.Get()
		childName := childCLINode.name

		childCmd, set := childCLINode.cmd.Get()
		if !set {
			childCmd = &cobra.Command{}
		}

		if childCmd.Use == "" {
			childCmd.Use = childName
		}

		childCLINode.cmd = mo.Some(childCmd)
		childTreeNode.Set(childCLINode)

		cmd.AddCommand(childCmd)

		c.walk(childTreeNode)
	}
}
