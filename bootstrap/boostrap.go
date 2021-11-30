package bootstrap

import (
	config "blog/config"
)

func Initialization() {
	//初始化配置
	config.Initialization()
	//连接数据库
	SetupDB()
}