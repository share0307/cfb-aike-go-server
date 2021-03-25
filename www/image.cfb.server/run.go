package main

import "image-cfb-server/command"

/**
	执行方法
 */
func main()  {

	// 初始化环境配置，可以不在此处初始化，放到相关的init方法里




	// 执行命令
	command.Initialize();
	command.Execute();

}
