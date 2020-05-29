package persistent

import (
    "../model"
    "../errors"
    "encoding/json"
    "io/ioutil"
    "log"
    "os"
)

const maxLevel int = 10
const pathName string = "storage/"

type CompanyFilePersistent struct {
}

func (c CompanyFilePersistent) save(data []model.Company) error {
    path, err := getConfigFilePath(pathName, maxLevel)
    if nil != err {
        log.Fatal("保存数据到文件，未找到配置路径: ", pathName)
        return errors.StockDataError{}
    }
    marshal, err := json.Marshal(data)
    if nil != err {
        log.Fatal("保存数据到文件，数据格式化异常: ", pathName)
        return errors.StockDataError{}
    }
    ioutil.WriteFile(path+getFileName(), marshal, os.ModeAppend)
    if nil != err {
        log.Fatal("寻找配置文件，未找到配置路径: ", pathName)
        return errors.StockDataError{}
    }
    return nil
}

func (c CompanyFilePersistent) read() []model.Company {
    path, err := getConfigFilePath(pathName, maxLevel)
    if nil != err {
        log.Fatal("从文件读取数据，未找到配置路径: ", pathName)
    }
    file, err := ioutil.ReadFile(path + getFileName())
    if nil != err {
        log.Fatal("从文件读取数据，读取数据异常: ", pathName)
    }

    d := &[]model.Company{}
    err = json.Unmarshal(file, &d)
    if nil != err {
        log.Fatal("从文件读取数据,解析数据异常: ", pathName)
    }
    return *d
}

func getConfigFilePath(pathName string, level int) (string, error) {
    log.Println("寻找配置文件，配置文件路径: ", pathName, " 向上遍历层级: ", maxLevel-level)
    if level < 0 {
        log.Fatal("寻找配置文件，配置文件路径: ", pathName)
        return "", errors.StockDataError{"未找到配置文件"}
    }
    _, err := os.Open(pathName)
    if nil == err {
        return pathName, nil
    } else {
        return getConfigFilePath("../"+pathName, level-1)
    }
}

func getFileName() string {
    return "company.json"
}
