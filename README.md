
1) 修改了参数的默认值

2) 新加参数ipPrefix，可以制定ip前缀，registrator会找以这个前缀开始的ip，比如 -ipPrefix=168.168，就会找168.168.×.×的ip

3) 修改了etcd的key、value格式，key的格式：/v2/keys/启动时指定的namespace/image的名字/container的id，value的格式：outerIp,outerPort:innerPort,outerPort:innerPort ...

4) docker内可以通过该方式获取container id （cat /proc/self/cgroup | grep -o  -e "docker-.*.scope" | head -n 1 | sed "s/docker-\(.*\).scope/\\1/"）

5) 因为启动时的namespace、image名字等都是提前知道的，container内部又能获得containerID，所以通过etcd查看，就可以知道这个container对外的ip和端口映射

6) 启动方式：DOCKER_HOST=unix:///var/run/docker.sock registrator -ipPrefix=168.168 etcd://localhost:2379/services

注：只支持etcd
