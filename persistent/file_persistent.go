package persistent

import (
    "../errors"
    "../model"
    "../s-logger"
    "../tool"
    "encoding/json"
    "io/ioutil"
    "os"
)

const (
    maxLevel int    = 10
    pathName string = "storage/"
)

var logger = s_logger.New()

type CompanyFilePreserver struct{}

func (c CompanyFilePreserver) Save(data []model.Company) error {
    path, err := tool.GetPath(pathName, maxLevel, maxLevel)
    if nil != err {
        logger.Infow("保存数据到文件,未找到配置路径", "pathName", pathName, "err", err)
        return errors.StockDataError{Msg: "保存数据到文件,未找到配置路径"}
    }
    marshal, err := json.MarshalIndent(data, "", "  ")
    if nil != err {
        logger.Infow("保存数据到文件,数据格式化异常", "pathName", pathName, "err", err)
        return errors.StockDataError{Msg: "保存数据到文件,数据格式化异常"}
    }
    err = ioutil.WriteFile(path+c.getFullFileName(tool.NowDate()), marshal, os.ModeAppend)
    if nil != err {
        logger.Infow("保存数据到文件,写入文件数据异常", "pathName", pathName, "err", err)
        return errors.StockDataError{Msg: "保存数据到文件,写入文件数据异常"}
    }
    return nil
}

func (c CompanyFilePreserver) Read() ([]model.Company, error) {
    path, err := tool.GetPath(pathName, maxLevel, maxLevel)
    if nil != err {
        logger.Infow("从文件读取数据,未找到配置路径", "pathName", pathName, "err", err)
        return nil, errors.StockDataError{Msg: "从文件读取数据,未找到配置路径"}
    }
    file, err := ioutil.ReadFile(path + c.getFullFileName(tool.NowDate()))
    if nil != err {
        logger.Infow("从文件读取数据,读取数据异常", "pathName", pathName, "err", err)
        return nil, errors.StockDataError{Msg: "从文件读取数据,读取数据异常"}
    }

    d := &[]model.Company{}
    err = json.Unmarshal(file, &d)
    if nil != err {
        logger.Infow("从文件读取数据,解析数据异常", "pathName", pathName, "err", err)
        return nil, errors.StockDataError{Msg: "从文件读取数据,读取数据异常"}
    }
    return *d, nil
}

func (c CompanyFilePreserver) getFullFileName(append string) string {
    return c.getPrefix() + append + c.getSuffix()
}

func (c CompanyFilePreserver) getPrefix() string {
    return "company-"
}

func (c CompanyFilePreserver) getSuffix() string {
    return ".json"
}
