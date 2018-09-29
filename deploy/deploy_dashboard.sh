#!/bin/bash
kubectl config use-context docker-for-desktop && kubectl cluster-info &> /dev/null
if [[ $? == 0 ]]; then
    echo "cluster start"
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/master/src/deploy/recommended/kubernetes-dashboard.yaml 
    kubectl proxy & &> /dev/null
    sleep 1
    open "http://localhost:8001/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/"
else
    echo "cluster shotdown"
    exit 1
fi
