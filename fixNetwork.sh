#没有网络图标 且 连不上网
# 停止服务
service NetworkManager stop
# 删除缓冲文件
sudo rm /var/lib/NetworkManager/NetworkManager.state
# 启动服务
service NetworkManager start