package cluster

import (
	_ "embed"

	"cks-cli/pkg/commands/utils"

	"github.com/urfave/cli/v2"
)

//go:embed scripts/create-token.sh
var createTokenScript string

//go:embed scripts/join.sh
var joinScript string

// ImageRun runs scan on docker image
func CreateToken(ctx *cli.Context) error {
	values := make(map[string]interface{})
	// values["version"] = ctx.String("version")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(createTokenScript, values))
}

// cks cluster join \
//  --master 192.168.0.8:6443
//  --token 0ns3xh.z8ts84ld79kf4pha
//  --discovery-token-ca-cert-hash sha256:e049e4827e20734500848a8f56c6043069b3b8650533cd590413af61426a7d87
func Join(ctx *cli.Context) error {
	values := make(map[string]interface{})
	values["master"] = ctx.String("master")
	values["token"] = ctx.String("token")
	values["cahash"] = ctx.String("ca-hash")
	// fmt.Println(s)
	// fmt.Println(ProcessString(s, values))
	return utils.RunCommand("/bin/sh", "-c", utils.ProcessString(joinScript, values))
}
