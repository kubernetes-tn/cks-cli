#!/bin/sh
echo "### Execute the command below in Worker Nodes ###"
kubeadm token create --print-join-command --ttl 0 | sed 's/kubeadm join/cks cluster join --master/g' | sed 's/discovery-token-ca-cert-hash/ca-hash/g'
