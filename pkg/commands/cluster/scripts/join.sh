#!/bin/sh
# cks cluster join --master {{ .master }} --token {{ .token }} --ca-hash {{ .cahash }}
kubeadm join {{ .master }} --token {{ .token }} --discovery-token-ca-cert-hash {{ .cahash }}