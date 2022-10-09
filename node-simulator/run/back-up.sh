kubebuilder create api --group infra --version v1 --kind KeepJob

$GOPATH/src/k8s.io/code-generator/generate-groups.sh all node-simulator/generated/keepjob node-simulator/api keepjob:v1 \
&& mv ~/go/src/node-simulator/generated/keepjob/ generated/