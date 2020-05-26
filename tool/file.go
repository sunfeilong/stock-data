package tool

import (
    "errors"
    "log"
    "os"
)

func GetConfigFilePath(fileName string, level int, maxLevel int) (string, error) {
    log.Println("寻找配置文件，配置文件路径: ", fileName, " 向上遍历层级: ", maxLevel-level)
    _, err := os.Open(fileName)
    if level < 0 {
        log.Println("寻找配置文件，配置文件路径: ", fileName)
        return "", errors.New("遍历目录未找到配置文件")
    }
    if nil == err {
        return fileName, nil
    } else {
        return GetConfigFilePath("../"+fileName, level-1, maxLevel)
    }
}
