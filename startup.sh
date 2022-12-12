 minikube start
 kubectl apply -f https://raw.githubusercontent.com/volcano-sh/volcano/master/installer/volcano-development.yaml
# ====================================  启动 colocation_backend  =========================================
cd colocation_backend
nohup bee run &
cd ..
# ====================================  启动 mockserver  =========================================
cd data_processor/mock-server/run
kubectl apply -f deployment.yaml
cd -
# ====================================  启动 node-simulator  =========================================
cd node-simulator
make manifests
make install
make run
cd -