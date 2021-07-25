# Quick Start

## Cluster Setup

**install master**

Identify your master machine and run

```
$ cks cluster install-master

```

<details>
<summary>Result</summary>

```
OS type is checked
Reading package lists...
Building dependency tree...
Reading state information...
bash-completion is already the newest version (1:2.10-1ubuntu1).
binutils is already the newest version (2.34-6ubuntu1.1).
0 upgraded, 0 newly installed, 0 to remove and 65 not upgraded.
Reading package lists...
Building dependency tree...
Reading state information...
Package 'docker-ce' is not installed, so not removed
Package 'docker-engine-cs' is not installed, so not removed
Package 'docker' is not installed, so not removed
Package 'docker-compose' is not installed, so not removed
Package 'docker-registry' is not installed, so not removed
Package 'docker2aci' is not installed, so not removed
Package 'docker-doc' is not installed, so not removed
Package 'docker-engine' is not installed, so not removed
The following packages were automatically installed and are no longer required:
  bridge-utils conntrack cri-tools ebtables pigz runc socat ubuntu-fan
Use 'apt autoremove' to remove them.
The following packages will be REMOVED:
  containerd docker.io kubeadm kubectl kubelet kubernetes-cni
.....
....
....
### Execute the command below in Worker Nodes ###
cks cluster join --master 172.31.118.222:6443 --token ppblkq.6uafwx1q03m0cxbq --ca-hash sha256:03466f37be5072fa68d84f156a1f68ce2e19d2e3f24833674263fa65b724acf9
```

</details>


**Install worker node**

Identify your worker machine(s) and run

```
$ cks cluster install-worker

```

<details>
<summary>Result</summary>

```
Reading package lists...
Building dependency tree...
Reading state information...
bash-completion is already the newest version (1:2.10-1ubuntu1).
binutils is already the newest version (2.34-6ubuntu1.1).
0 upgraded, 0 newly installed, 0 to remove and 26 not upgraded.
Reading package lists...
Building dependency tree...
Reading state information...
Package 'docker-engine-cs' is not installed, so not removed
Package 'docker' is not installed, so not removed
Package 'docker-compose' is not installed, so not removed
Package 'docker-registry' is not installed, so not removed
Package 'docker2aci' is not installed, so not removed
Package 'docker-doc' is not installed, so not removed
Package 'containerd.io' is not installed, so not removed
Package 'docker-ce-cli' is not installed, so not removed
Package 'docker-ce-rootless-extras' is not installed, so not removed
Package 'docker-ce' is not installed, so not removed
Package 'docker-scan-plugin' is not installed, so not removed
Package 'docker-engine' is not installed, so not removed
The following packages were automatically installed and are no longer required:
  bridge-utils conntrack cri-tools ebtables pigz runc socat ubuntu-fan
Use 'apt autoremove' to remove them.
.....
....
...
```

</details>


**worker joins a master**

Join the master ( run it in worker(s))
```
$ cks cluster join --master <master-host>:<api-port> --token <token> --ca-hash <sha256:xxxx>

```

example:

```
cks cluster join --master 172.31.118.222:6443 --token ppblkq.6uafwx1q03m0cxbq --ca-hash sha256:03466f37be5072fa68d84f156a1f68ce2e19d2e3f24833674263fa65b724acf9
```

<details>
<summary>Result</summary>
```
cks cluster join --master 172.31.118.222:6443 --token ppblkq.6uafwx1q03m0cxbq --ca-hash sha256:03466f37be5072fa68d84f156a1f68ce2e19d2e3f24833674263fa65b724acf9
[preflight] Running pre-flight checks
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...
This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.
Run 'kubectl get nodes' on the control-plane to see this node join the cluster
```
</details>

## Falco

**install falco**

Identify one of your worker nodes and install falco

```
$ cks falco install

```

<details>
<summary>Result</summary>

```
...
....
....

```

</details>


**Follow falco tutorial**

Enjoy a falco tutorial step by step

```
$ cks falco tuto

```

<details>
<summary>Result</summary>

```
 ____    _    ____ _____ 
/ ___|  / \  |  _ \_   _|
\___ \ / _ \ | |_) || |  
 ___) / ___ \|  _ < | |  
|____/_/   \_\_| \_\|_|
### Create Namespace for the demo ###

kubectl create ns demo-falco

# Are you done? [y/n]: 
......
.....
....
 _____ _   _ ____  
| ____| \ | |  _ \ 
|  _| |  \| | | | |
| |___| |\  | |_| |
|_____|_| \_|____/ 
### Clean up namespace

kubectl delete ns demo-falco


## check another advanced tuto in this link :
https://k8s.tn/rdiL4i

# Clean up done? [y/n]: y
END of Falco Tuto - Congrats!

```

</details>
