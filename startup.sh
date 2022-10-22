 minikube start
# ====================================  启动 colocation_backend  =========================================
cd colocation_backend
nohup bee run &
cd ..
# ====================================  启动 mockserver  =========================================
cd data_processor/mock-server/run
kubectl apply -f deployment
cd -
# ====================================  启动 node-simulator  =========================================
cd node-simulator
make manifests
make install
make deploy
cd -