package cluster

import (
	_ "embed"

	"cks-cli/pkg/commands/utils"

	"github.com/urfave/cli/v2"
)

//go:embed scripts/check-requirements.sh
var checkRequirementsScript string

//go:embed scripts/install-master.sh
var installMasterScript string

//go:embed scripts/install-worker.sh
var installWorkerScript string

func CheckRequirements(ctx *cli.Context) error {
	values := make(map[string]interface{})
	// values["version"] = ctx.String("version")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(checkRequirementsScript, values))
}

func InstallMaster(ctx *cli.Context) error {
	values := make(map[string]interface{})
	values["version"] = ctx.String("version")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(installMasterScript, values))
}

func InstallWorker(ctx *cli.Context) error {
	values := make(map[string]interface{})
	values["version"] = ctx.String("version")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(installWorkerScript, values))
}
