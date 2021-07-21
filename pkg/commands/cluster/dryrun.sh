KUBE_VERSION={{ .version }};
echo "Installing k8s master $KUBE_VERSION"
echo apt-get remove -y docker.io kubelet kubeadm kubectl kubernetes-cni
sleep 4
echo apt-get autoremove -y
sleep 2
echo apt-get install -y etcd-client vim build-essential
sleep 4