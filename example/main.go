package main

import (
	"os"

	cli "github.com/rancher/wrangler-cli"
	"github.com/rancher/wrangler-cli/example/pkg/app"
)

func main() {
	os.Setenv("MY_BOOL", "1")
	os.Setenv("MY_INT", "123")
	os.Setenv("MY_STRING", "hello")
	os.Args = append(os.Args, "sub")
	cli.Main(app.New())
}
