
1) 修改了参数的默认值

2) 修改了etcd key的格式，格式为：/v2/keys/启动时指定的前缀/image的名字/端口/container的id，value还是原来的hostIP:hostPort

3) docker内可以通过该方式获取container id （cat /proc/self/cgroup | grep -o  -e "docker-.*.scope" | head -n 1 | sed "s/docker-\(.*\).scope/\\1/"）

4) 因为启动时指定的前缀、image名字、端口都是提前知道的，所以查询etcd，一下就可以拿到某个端口的所有服务，又因为container内部也能获得container id的，这样也就知道了自己对外开放的ip和端口
