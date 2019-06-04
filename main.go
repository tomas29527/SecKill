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
		Endpoints:   []string{"106.13.60.183:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()

	if err != nil {
		logs.Debug("链接etcd错误: %v", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Put(ctx, "/app_server/server0", "192.1.1.3")
	cancel()
	if err != nil {
		logs.Debug("放数据失败: %v", err)
		return
	}
	logs.Debug("resp: %v", resp)

	//ctx1, cancel1 := context.WithTimeout(context.Background(), 2*time.Second)
	//response, err := cli.Get(ctx1, "/app_server/server0")
	//cancel1()
	//if err != nil {
	//	logs.Debug("拿数据失败: %v",err)
	//	return
	//}
	//logs.Debug("response: %v",response.Kvs)

	//beego.Run()
}
