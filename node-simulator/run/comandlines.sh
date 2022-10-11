$GOPATH/src/k8s.io/code-generator/generate-groups.sh all node-simulator/generated node-simulator/apis infra:v1
$GOPATH/src/k8s.io/code-generator/generate-groups.sh all colocation_backend/generated colocation_backend/apis infra:v1
$GOPATH/src/k8s.io/code-generator/generate-groups.sh all keep-resources/pkg/generated keep-resources/pkg/apis infra:v1
rm -rf generated/*
cp -r  ~/go/src/colocation_backend/generated/* generated/