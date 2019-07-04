package main

import (
	"SecKill/conf"
	"context"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	//初始化日志
	conf.LogSetting()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.EtcdConfObj.EtcdServer,
		DialTimeout: time.Duration(conf.EtcdConfObj.EtcdDialTimeout) * time.Second,
	})
	defer cli.Close()
	if err != nil {
		logs.Debug("链接etcd错误: %v", err)
		return
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//resp, err := cli.Put(ctx, "/app_server/server0", "192.1.1.3")
	//cancel()
	//if err != nil {
	//	logs.Debug("放数据失败: %v", err)
	//	return
	//}
	//logs.Debug("resp: %v", resp)

	//ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	//response, err := cli.Get(ctx1, "/app_server/server0")
	//cancel1()
	//if err != nil {
	//	logs.Debug("拿数据失败: %v",err)
	//	return
	//}
	//logs.Debug("response: %v",response.Kvs)
	go func() {
		watch := cli.Watch(context.Background(), "/app_server/server0", clientv3.WithPrefix())
		for wresp := range watch {
			for _, ev := range wresp.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					logs.Debug("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

				case clientv3.EventTypeDelete:
					logs.Debug("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				}
			}
		}
	}()

	//beego.Run()
}
