package falco

import (
	_ "embed"

	"cks-cli/pkg/commands/cluster"
	"cks-cli/pkg/commands/utils"

	"github.com/urfave/cli/v2"
)

//go:embed scripts/install.sh
var installScript string

func Install(ctx *cli.Context) error {
	err := cluster.CheckRequirements(ctx)
	if err != nil {
		return err
	}
	values := make(map[string]interface{})
	// values["version"] = ctx.String("version")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(installScript, values))
}
