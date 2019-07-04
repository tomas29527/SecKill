package init

import "go.etcd.io/etcd/clientv3"

type etcdServer struct {
	Name   string
	stop   chan bool
	client *clientv3.Client
}
