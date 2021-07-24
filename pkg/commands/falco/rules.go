package falco

import (
	_ "embed"
	"fmt"

	"github.com/urfave/cli/v2"
)

//go:embed files/rules.txt
var rules string

func Rules(ctx *cli.Context) error {

	fmt.Println(rules)
	return nil
}
