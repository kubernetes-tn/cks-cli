
# Overview

A comprehensive CLI for CKS exam preparation

# Getting Started

```sh
docker-compose run --rm cks help
```

# Prerequisites

- Create a set of virtual machines (Amazon EC2 or Azure VMs)
- One of the machines will be your master node and the other(s) will be your worker node(s)
- Note: Ensure the virtual machines have network connectivity between them otherwise you will not be able to establish a connection or joined the nodes together

# Installation of binaries

## Linux

```sh
- curl -o cks https://k8s.tn/cks-cli/latest/cks-linux-amd64
- chmod +x cks
- sudo mv cks /usr/local/bin
```

## OS X

```sh
- curl -o cks https://k8s.tn/cks-cli/latest/cks-darwin-amd64
- chmod +x cks
- sudo mv cks /usr/local/bin
```

# Cluster Setup

## Installing the master node

```sh
cks cluster install-master
```

once above is complete, you will get a command that will be executed on the worker node looking like:

```sh
cks cluster join --master <master-host>:<api-port> --token <token> --ca-hash <sha256:xxxx
```

## Installing the worker node

```sh
cks cluster install-worker
```
## Joining the worker node to the master 

```sh
cks cluster join --master <master-host>:<api-port> --token <token> --ca-hash <sha256:xxxx>
```

# Installing Falco on CKS CLI

```sh
cks falco install
```
# Now check your nodes are running 

- on your master node, run:

```sh
kubectl get nodes
```

You should now see your nodes running.
# Docs

https://cks.k8s.tn

# Authors

- @abdennour - [Abdennour TOUMI](https://www.linkedin.com/in/abdennour/)
- @moabukar - [Mohamed Abukar](https://linkedin.com/in/mohamed-abukar)

# License

[LICENSE](LICENSE)