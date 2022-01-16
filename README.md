# go_lear

#docker registry 一般使用内部仓库 这里就不写了。 docker push  XXXX:v0.0.1 开发权限即可

#本地镜像编辑：docker build -t golear_test:v0.0.1 .

#docker 利用nsenter查看ip信息

#docker_pid=$(docker ps   ｜ grep golear_test | awk '{print $1}')
# nsenter -t ($docker inspect $docker_pid | grep -i pid) -n ip a
