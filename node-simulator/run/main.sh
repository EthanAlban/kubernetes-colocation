#!/bin/bash
#ls
#cd ..
#echo "start compile"
#go build main.go
echo "10.11.176.109 cls-ggyf5dq3.ccs.tencent-cloud.com" >> /etc/hosts
echo "10.11.40.29 cls-iv3dg0fl.ccs.tencent-cloud.com" >> /etc/hosts

echo "start running"
#./go-monitor --kubeconfig=conf/kubeconfigs/kubeflow-config-cls-iv3dg0fl.yml --qs_env=production
./mock-server
echo "echo fatal err has happened,stay alive for 5 min for debugging"
sleep 5m
#curl -X GET "http://127.0.0.1:5000/api/monitor/mlsys/trials/calca_trial_gpu_hours?job_id=6&trial_id=8" -H  "accept: application/json"