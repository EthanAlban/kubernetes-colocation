::打包镜像
docker build -t reg.htres.cn/yw_18kn1/demonginx .
 
::推镜像
docker push reg.htres.cn/yw_18kn1/demonginx
 
::展示镜像
docker images
 
pause