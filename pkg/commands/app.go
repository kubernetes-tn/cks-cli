package commands

import (
	"cks-cli/commands/cluster"

	"github.com/urfave/cli/v2"
)

var (
	verisonFlag = cli.StringFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Value:   "1.21.0",
		Usage:   "k8s version",
		EnvVars: []string{"CKS_VERSION"},
	}

	typeFlag = cli.StringFlag{
		Name:    "type",
		Aliases: []string{"t"},
		Value:   "master",
		Usage:   "machine type( master or worker)",
		EnvVars: []string{"CKS_type"},
	}

	clusterFlags = []cli.Flag{
		&verisonFlag,
		&typeFlag,
	}
)

func NewApp(version string) *cli.App {
	// cli.VersionPrinter = func(c *cli.Context) {
	// 	showVersion(c.String("cache-dir"), c.String("format"), c.App.Version, c.App.Writer)
	// }

	app := cli.NewApp()
	app.Name = "cks"
	app.Version = version
	app.ArgsUsage = "target"
	app.Usage = "A CKS exam preparation command line"
	app.EnableBashCompletion = true
	// flags := append(globalFlags, setHidden(deprecatedFlags, true)...)
	// flags = append(flags, setHidden(imageFlags, true)...)

	// app.Flags = flags
	app.Commands = []*cli.Command{
		NewClusterCommand(),
	}

	return app
}

// NewImageCommand is the factory method to add image command
func NewClusterCommand() *cli.Command {
	return &cli.Command{
		Name:      "cluster",
		Aliases:   []string{"c"},
		ArgsUsage: "cluster_install",
		Usage:     "manage cluster",
		Subcommands: cli.Commands{
			{
				Name:      "install-master",
				Aliases:   []string{"im"},
				Usage:     "install k8s on a machine",
				ArgsUsage: "",
				Action:    cluster.InstallMaster,
				Flags:     clusterFlags,
			},
		},
	}
}
