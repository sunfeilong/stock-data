package persistent

import (
    "encoding/json"
    "errors"
    "github.com/xiaotian/stock/pkg/model"
    "github.com/xiaotian/stock/pkg/s-logger"
    "github.com/xiaotian/stock/pkg/tool"
    "io/ioutil"
    "os"
)

var logger = s_logger.New()

type CompanyFilePreserver struct{}

func (c CompanyFilePreserver) Save(data []model.Company) error {
    path, err := tool.GetPath(pathName, maxLevel, maxLevel)
    if nil != err {
        logger.Infow("保存数据到文件,未找到配置路径", "pathName", pathName, "err", err)
        return errors.New("保存数据到文件,未找到配置路径")
    }
    marshal, err := json.Marshal(data)
    if nil != err {
        logger.Infow("保存数据到文件,数据格式化异常", "pathName", pathName, "err", err)
        return errors.New("保存数据到文件,数据格式化异常")
    }
    err = ioutil.WriteFile(path+c.getFullFileName(tool.NowDate()), marshal, os.ModeAppend)
    if nil != err {
        logger.Infow("保存数据到文件,写入文件数据异常", "pathName", pathName, "err", err)
        return errors.New("保存数据到文件,写入文件数据异常")
    }
    return nil
}

func (c CompanyFilePreserver) Read() ([]model.Company, error) {
    path, err := tool.GetPath(pathName, maxLevel, maxLevel)
    if nil != err {
        logger.Infow("从文件读取数据,未找到配置路径", "pathName", pathName, "err", err)
        return nil, errors.New("从文件读取数据,未找到配置路径")
    }
    file, err := ioutil.ReadFile(path + c.getFullFileName(tool.NowDate()))
    if nil != err {
        logger.Infow("从文件读取数据,读取数据异常", "pathName", pathName, "err", err)
        return nil, errors.New("从文件读取数据,读取数据异常")
    }

    d := &[]model.Company{}
    err = json.Unmarshal(file, &d)
    if nil != err {
        logger.Infow("从文件读取数据,解析数据异常", "pathName", pathName, "err", err)
        return nil, errors.New("从文件读取数据,读取数据异常")
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
