#!/usr/bin/env bash

kubectl delete pods -A --all --force --grace-period=0 --wait=false
kubectl delete services -A --all --force --grace-period=0 --wait=false
kubectl delete endpoints -A --all --force --grace-period=0 --wait=false
echo '{"kind":"Namespace","spec":{"finalizers":[]},"apiVersion":"v1","metadata":{"name":"workload"}}' | kubectl replace --raw '/api/v1/namespaces/workload/finalize' -f -
