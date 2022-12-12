# directly run this script to generate new generated folders and copy to node-simulator
echo "====================  update clienset  ====================="
cd /home/et/Desktop/Projects/keep_colocation/keep-resources
cp -r pkg/apis/ ~/go/src/github.com/keep-resources/pkg/
$GOPATH/src/k8s.io/code-generator/generate-groups.sh all github.com/keep-resources/pkg/generated github.com/keep-resources/pkg/apis infra:v1
cd /home/et/Desktop/Projects/keep_colocation/keep-resources/pkg
cp -r ~/go/src/github.com/keep-resources/pkg/generated/ .
echo "==================  copy to no-simulator  =================="
rm -rf /home/et/Desktop/Projects/keep_colocation/node-simulator/api/infra/v1
cp -r /home/et/Desktop/Projects/keep_colocation/keep-resources/pkg/apis/infra/v1 /home/et/Desktop/Projects/keep_colocation/node-simulator/api/infra
cd /home/et/Desktop/Projects/keep_colocation/node-simulator/
echo "=======================  install crd  ======================"
make uninstall
make manifests
make install
echo "=======================   completed   ======================"