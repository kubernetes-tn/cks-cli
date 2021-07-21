package cluster

import (
	_ "embed"

	"cks-cli/pkg/commands/utils"

	"github.com/urfave/cli/v2"
)

//go:embed dryrun.sh
var installMasterScript string

// ImageRun runs scan on docker image
func InstallMaster(ctx *cli.Context) error {
	values := make(map[string]interface{})
	values["version"] = ctx.String("version")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(installMasterScript, values))
}
