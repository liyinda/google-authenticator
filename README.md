# security_exporter

Simple server that scrapes security stats and exports them via HTTP for Prometheus consumption

To support time related histogram metrics, Used to check hacker invason such as reverse Shell.

基于Prometheus监控体系，对主机的安全性进行监控，探测主机是否存在反弹shell，多次ssh尝试登陆等迹象。

## Table of Contents
* [Dependency](#dependency)
* [Download](#download)
* [Compile](#compile)
  * [build binary](#build-binary)
  * [build docker image](#build-docker-image)
* [Run](#run)
  * [run binary](#run-binary)
  * [run docker image](#run-docker-image)
* [Environment variables](#environment-variables)
* [Metrics](#metrics)
  * [Server main](#server-main)
  * [Server zones](#server-zones)
  * [Filter zones](#filter-zones)



## Dependency

* [lsof](http://www.linuxfromscratch.org/blfs/view/svn/general/lsof.html)
* [Prometheus](https://prometheus.io/)
* [Golang 1.9.4](https://golang.org/)


## Download

Binary can be downloaded from [Releases](https://github.com/liyinda/secuity_exporter/releases) page.

## Compile

### build binary

``` shell
go build security_exporter.go
```
### build docker image
``` shell
make docker
```

## Docker Hub Image
``` shell
DOCKER 部署方式作者会尽快补充 
docker pull 空:latest
```
### run docker
```
docker run  -ti 镜像地址 bin/security_exporter
```

## Environment variables

This image is configurable using different env variables

## Metrics

Documents about exposed Prometheus metrics.

``` 
# HELP fail_password_total Number of Fail Password in /var/log/secure.
# TYPE fail_password_total counter
fail_password_total{host="$hostname",zone="datacenter"} 3
# HELP file_change_total Number of Change in /etc.
# TYPE file_change_total counter
file_change_total{host="$hostname",zone="datacenter"} 21
# HELP reverse_shell_total Number of Reverse Shell.
# TYPE reverse_shell_total counter
reverse_shell_total{host="$hostname",zone="datacenter"} 0

```

### Grafana

![image](https://github.com/liyinda/security_exporter/blob/master/jpg/grafana.jpg)


