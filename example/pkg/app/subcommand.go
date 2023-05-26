package app

import (
	"fmt"

	"github.com/spf13/cobra"

	cli "github.com/rancher/wrangler-cli"
)

type private struct{}

type Embedded struct {
	MyBool   bool   `name:"my-bool" env:"MY_BOOL"`
	MyInt    int    `name:"my-int" env:"MY_INT"`
	MyString string `name:"my-string" env:"MY_STRING" default:"noop"`
}

func NewSubCommand() *cobra.Command {
	return cli.Command(&SubCommand{}, cobra.Command{
		Short: "Add some short description",
		Long:  "Add some long description",
	})
}

type SubCommand struct {
	Embedded
	OptionOne string `usage:"Some usage description"`
	OptionTwo string `name:"custom-name"`

	privateField private
}

func (s *SubCommand) Run(cmd *cobra.Command, args []string) error {
	fmt.Println("I do stuff with env bool:", s.MyBool, "and env int:", s.MyInt, "and env string:", s.MyString)
	return nil
}
