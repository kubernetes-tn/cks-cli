# Cluster

```bash
NAME:
   cks cluster - manage cluster (support Ubuntu 20 OS only for now)

USAGE:
   cks cluster command [command options] cluster_install

COMMANDS:
   check-requirements, req  check if your machines meet requirements to install the cluster
   install-master, im       install k8s master components
   install-worker, iw       install k8s worker components
   create-token, ct         Generate a token and print the command of joining a master k8s. It must run inside a master machine
   join, j                  Join a kubernetes master. Must be run in a worker node
   help, h                  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```