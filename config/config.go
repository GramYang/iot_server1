package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var Conf =&Config{}

type Config struct{
	ServerPort string `json:"server_port"`
	MysqlAddr     string `json:"mysql_addr"`
	MysqlPort     string `json:"mysql_port"`
	MysqlUserName string `json:"mysql_username"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`
	LocalDebug bool `json:"local_debug"`
}

func(c *Config) defaultConfig() {
	c.ServerPort=":8086"
	c.MysqlAddr="106.54.87.204"
	c.MysqlPort="3306"
	c.MysqlUserName="cqdq"
	c.MysqlPassword="cqdq12345"
	c.MysqlDatabase="iot_admin"
	c.LocalDebug=true
}

func(c *Config) initConfig(path string){
	c.defaultConfig()
	if path!=""{
		file,err:=ioutil.ReadFile(path)
		if err!=nil{
		}
		if err=json.Unmarshal(file,c);err!=nil{
		}
	}
}

func Setup(){
	var p string
	flag.StringVar(&p,"c","","配置文件路径")
	flag.Parse()
	Conf.initConfig(p)
}