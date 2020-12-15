package tool

import (
    "errors"
    "github.com/xiaotian/stock/pkg/s-logger"
    "os"
)

var logger = s_logger.New()

func GetConfigFilePath(fileName string, level int, maxLevel int) (string, error) {
    logger.Infow("查找文件,配置文件路径", "fileName", fileName, "up level", maxLevel-level)
    _, err := os.Open(fileName)
    if level < 0 {
        logger.Infow("查找文件,未找到配置文件", "fileName", fileName, "up level", maxLevel-level)
        return "", errors.New("查找文件,未找到配置文件")
    }
    if nil == err {
        return fileName, nil
    } else {
        return GetConfigFilePath("../"+fileName, level-1, maxLevel)
    }
}

func GetPath(pathName string, level int, maxLevel int) (string, error) {
    logger.Infow("查找文件,配置文件路径", "fileName", pathName, "up level", maxLevel-level)
    if level < 0 {
        logger.Infow("查找文件,未找到配置文件", "fileName", pathName, "up level", maxLevel-level)
        return "", errors.New("查找文件,未找到配置文件")
    }
    _, err := os.Open(pathName)
    if nil == err {
        return pathName, nil
    } else {
        return GetPath("../"+pathName, level-1, maxLevel)
    }
}
