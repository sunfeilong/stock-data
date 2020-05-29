package main

import (
    "./s-logger"
)

var logger = s_logger.New()

func main() {

    logger.Infow("项目启动")

    logger.Infow("项目运行结束")

}
