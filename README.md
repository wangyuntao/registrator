
1) 修改了参数的默认值

2) 新加参数ipPrefix，可以制定ip前缀，registrator会找以这个前缀开始的ip，比如 -ipPrefix=168.168，就会找168.168.×.×的ip

3) 修改了etcd key的格式，格式为：/v2/keys/启动时指定的前缀/image的名字/端口/container的id，value还是原来的hostIP:hostPort

4) docker内可以通过该方式获取container id （cat /proc/self/cgroup | grep -o  -e "docker-.*.scope" | head -n 1 | sed "s/docker-\(.*\).scope/\\1/"）

5) 因为启动时指定的前缀、image名字、端口都是提前知道的，所以查询etcd，一下就可以拿到某个端口的所有服务，又因为container内部也能获得container id的，这样也就知道了自己对外开放的ip和端口

6) 启动方式：DOCKER_HOST=unix:///var/run/docker.sock registrator -ipPrefix=168.168 etcd://localhost:2379/services
