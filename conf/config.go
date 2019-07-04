package conf

import (
	"SecKill/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type EtcdConf struct {
	EtcdServer            []string
	EtcdDialTimeout       int
	EtcdSeckillPrefixKey  string
	EtcdSeckillProductKey string
}

var EtcdConfObj *EtcdConf

func initEtcdConf() error {
	server := beego.AppConfig.Strings("prod::etcd_server")
	logs.Debug("====server:%v", server)
	timeout, err := beego.AppConfig.Int("prod::etcd_dial_timeout")
	if err != nil {
		return err
	}
	EtcdConfObj = &EtcdConf{
		EtcdServer:      server,
		EtcdDialTimeout: timeout,
	}
	return nil
}

func init() {
	//etcd配置初始化
	err := initEtcdConf()
	if err != nil {
		panic("init etcd conf is err")
	}
}

// beego 日志配置结构体
type LoggerConfig struct {
	FileName            string   `json:"filename"`
	Level               int      `json:"level"`    // 日志保存的时候的级别，默认是 Trace 级别
	Maxlines            int      `json:"maxlines"` // 每个文件保存的最大行数，默认值 1000000
	Maxsize             int      `json:"maxsize"`  // 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	Daily               bool     `json:"daily"`    // 是否按照每天 logrotate，默认是 true
	Maxdays             int      `json:"maxdays"`  // 文件最多保存多少天，默认保存 7 天
	Rotate              bool     `json:"rotate"`   // 是否开启 logrotate，默认是 true
	Perm                string   `json:"perm"`     // 日志文件权限
	RotatePerm          string   `json:"rotateperm"`
	EnableFuncCallDepth bool     `json:"-"`        // 输出文件名和行号
	LogFuncCallDepth    int      `json:"-"`        // 函数调用层级
	Separate            []string `json:"separate"` //需要单独写入文件的日志级别,设置后命名类似 test.error.log
}

//日志运行级别检测变量
func LogSetting() {
	//获取运行方式
	loglevel := beego.AppConfig.String("runmode")
	if loglevel == "prod" {
		//处于生产模式,检测日志运行类型
		logtype := beego.AppConfig.String("prod::logtype")
		logname := beego.AppConfig.String("prod::logname")
		separate := []string{"notice", "error"}
		var logCfg = LoggerConfig{
			FileName:            "logs/" + logname,
			Level:               logs.LevelDebug,
			EnableFuncCallDepth: true,
			LogFuncCallDepth:    3,
			RotatePerm:          "777",
			Perm:                "777",
			Separate:            separate,
		}
		if logtype == "file" {
			err := utils.CheckAndDirCreate("logs")
			if err == nil {
				cfg, _ := json.Marshal(&logCfg)
				_ = beego.SetLogger(logs.AdapterMultiFile, string(cfg))
			}
		}
	}
}
