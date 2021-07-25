package commands

import (
	"cks-cli/pkg/commands/cluster"
	"cks-cli/pkg/commands/falco"

	"github.com/urfave/cli/v2"
)

var (
	topicFlag = cli.StringFlag{
		Name:    "topic",
		Aliases: []string{"t"},
		Value:   "",
		Usage:   "Topic from CKS learning objectives",
		EnvVars: []string{"CKS_TOPIC"},
	}
	verisonFlag = cli.StringFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Value:   "1.21.0",
		Usage:   "k8s version",
		EnvVars: []string{"CKS_VERSION"},
	}

	tokenFlag = cli.StringFlag{
		Name:    "token",
		Aliases: []string{"t"},
		Value:   "",
		Usage:   "token to join the master node",
		EnvVars: []string{"CKS_TOKEN"},
	}

	masterFlag = cli.StringFlag{
		Name:    "master",
		Aliases: []string{"m"},
		Value:   "",
		Usage:   "the endpoint of the master i.e. <hostname>:6443",
		EnvVars: []string{"CKS_MASTER"},
	}

	caHashFlag = cli.StringFlag{
		Name:    "ca-hash",
		Aliases: []string{"ch"},
		Value:   "",
		Usage:   "discovery token CA cert hash .i.e: sha256:e049e4827e20734500848a8f56c6043069b3b8650533cd590413af61426a7d87",
		EnvVars: []string{"CKS_CA_HASH"},
	}

	// Global flags
	globalFlags = []cli.Flag{
		&topicFlag,
	}

	clusterInstallFlags = []cli.Flag{
		&verisonFlag,
	}

	clusterJoinFlags = []cli.Flag{
		&masterFlag,
		&tokenFlag,
		&caHashFlag,
	}

	falcoInstallFlags = []cli.Flag{}
)

func NewApp(version string) *cli.App {
	// cli.VersionPrinter = func(c *cli.Context) {
	// 	showVersion(c.String("cache-dir"), c.String("format"), c.App.Version, c.App.Writer)
	// }

	app := cli.NewApp()
	app.Name = "cks"
	app.Version = version
	app.ArgsUsage = "verb"
	app.Usage = "A Comprehensive CLI for CKS exam preparation"
	app.EnableBashCompletion = true
	// flags := append(globalFlags, setHidden(deprecatedFlags, true)...)
	// flags = append(flags, setHidden(imageFlags, true)...)

	// app.Flags = flags
	app.Commands = []*cli.Command{
		NewClusterCommand(),
		NewFalcoCommand(),
	}

	return app
}

// NewClusterCommand is the factory method to add image command
func NewClusterCommand() *cli.Command {
	return &cli.Command{
		Name:      "cluster",
		Aliases:   []string{"c"},
		ArgsUsage: "cluster_install",
		Usage:     "manage cluster (support Ubuntu 20 OS only for now)",
		Subcommands: cli.Commands{
			{
				Name:      "check-requirements",
				Aliases:   []string{"req"},
				Usage:     "check if your machines meet requirements to install the cluster",
				ArgsUsage: "",
				Action:    cluster.CheckRequirements,
			},
			{
				Name:      "install-master",
				Aliases:   []string{"im"},
				Usage:     "install k8s master components",
				ArgsUsage: "",
				Action:    cluster.InstallMaster,
				Flags:     clusterInstallFlags,
			},
			{
				Name:      "install-worker",
				Aliases:   []string{"iw"},
				Usage:     "install k8s worker components",
				ArgsUsage: "",
				Action:    cluster.InstallWorker,
				Flags:     clusterInstallFlags,
			},
			{
				Name:      "create-token",
				Aliases:   []string{"ct"},
				Usage:     "Generate a token and print the command of joining a master k8s. It must run inside a master machine",
				ArgsUsage: "",
				Action:    cluster.CreateToken,
			},
			{
				Name:      "join",
				Aliases:   []string{"j"},
				Usage:     "Join a kubernetes master. Must be run in a worker node",
				ArgsUsage: "--master 192.168.0.8:6443 --token ztyj0p.cdew3ak071o57t7w --ca-hash sha256:e049e4827e20734500848a8f56c6043069b3b8650533cd590413af61426a7d8",
				Action:    cluster.Join,
				Flags:     clusterJoinFlags,
			},
		},
	}
}

func NewFalcoCommand() *cli.Command {
	return &cli.Command{
		Name:      "falco",
		Aliases:   []string{"f"},
		ArgsUsage: "falco_install",
		Usage:     "teach falco in CKS context",
		Subcommands: cli.Commands{
			{
				Name:      "check-requirements",
				Aliases:   []string{"req"},
				Usage:     "check if your machines meet requirements to install the cluster",
				ArgsUsage: "",
				Action:    cluster.CheckRequirements,
			},
			{
				Name:      "install",
				Aliases:   []string{"i"},
				Usage:     "install falco (support Ubuntu 20 OS only for now)",
				ArgsUsage: "",
				Action:    falco.Install,
				Flags:     falcoInstallFlags,
			},
			{
				Name:      "tutorial",
				Aliases:   []string{"t", "tuto"},
				Usage:     "tuto falco",
				ArgsUsage: "",
				Action:    falco.Tuto,
			},
			{
				Name:      "rules",
				Aliases:   []string{"r"},
				Usage:     "rules falco",
				ArgsUsage: "",
				Action:    falco.Rules,
			},
		}}
}
